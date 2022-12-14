# Licensed to Elasticsearch B.V. under one or more contributor
# license agreements. See the NOTICE file distributed with
# this work for additional information regarding copyright
# ownership. Elasticsearch B.V. licenses this file to you under
# the Apache License, Version 2.0 (the "License"); you may
# not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# 	http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

"""Constraints solver helper class."""

import random
from functools import wraps

from ..constraints import ConflictError

_max_attempts = 100000

ecs_constraints = {
    "bytes": [(">=", 0), ("<", 2**32)],
    "pid": [(">", 0), ("<", 2**32)],
    "port": [(">", 0), ("<", 2**16)],
}


def get_ecs_constraints(field):
    if field in ecs_constraints:
        return ecs_constraints[field]
    field = field.split(".")[-1]
    if field in ecs_constraints:
        return ecs_constraints[field]
    return []


def delete_by_cond(list, cond):
    for i in reversed([i for i, x in enumerate(list) if cond(x)]):
        del list[i]


def delete_use_once(list):
    def is_use_once(item):
        return len(item) > 2 and item[2] and item[2].get("use_once", False)

    delete_by_cond(list, is_use_once)


class solver:  # noqa: N801
    solvers = {}

    def __init__(self, name, *args):
        self.name = name
        self.valid_constraints = ("join_value", "max_attempts", "cardinality") + args

    def wrap_field_solver(self, func):
        @wraps(func)
        def _solver(field, value, constraints, environment):
            join_values = []
            max_attempts = None
            cardinality = 0
            history = []
            augmented_constraints = constraints + get_ecs_constraints(field)
            for k, v, *_ in augmented_constraints:
                if k not in self.valid_constraints:
                    raise NotImplementedError(f"Unsupported {self.name} constraint: {k}")
                if k == "join_value":
                    join_values.append(v)
                if k == "max_attempts":
                    v = int(v)
                    if v < 0:
                        raise ValueError(f"max_attempts cannot be negative: {v}")
                    if max_attempts is None or max_attempts > v:
                        max_attempts = v
                if k == "cardinality":
                    if type(v) is tuple:
                        if len(v) > 1:
                            raise ValueError(f"Too many arguments for cardinality of '{field}': {v}")
                        v = v[0]
                    cardinality = int(v)
                    history = environment.setdefault("fields_history", {}).setdefault(field, [])
            if max_attempts is None:
                max_attempts = _max_attempts
            if len(history) < cardinality:
                augmented_constraints.extend(("!=", v["value"]) for v in history)
            if not cardinality or len(history) < cardinality:
                value = func(field, value, augmented_constraints, max_attempts + 1, environment)
                if not value["left_attempts"]:
                    raise ConflictError(f"attempts exausted: {max_attempts}", field)
                del value["left_attempts"]
                if cardinality:
                    history.append(value)
            else:
                value = random.choice(history[:cardinality])
            for field, constraint in join_values:
                constraint.append_constraint(field, "==", value["value"], {"use_once": True})
            delete_use_once(constraints)
            return value

        return _solver

    def __call__(self, func):
        if not self.name.endswith("."):
            func = self.wrap_field_solver(func)
        self.solvers[self.name] = func
        return func

    @classmethod
    def solve_field(cls, field, constraints, schema, environment):
        if constraints is None:
            return None
        field_schema = schema.get(field, {})
        field_type = field_schema.get("type", "keyword")
        try:
            solver = cls.solvers[field_type]
        except KeyError:
            raise NotImplementedError(f"Constraints solver not implemented: {field_type}")
        if "array" in field_schema.get("normalize", []):
            value = []
        else:
            value = None
        return solver(field, value, constraints, environment)["value"]

    @classmethod
    def solve_nogroup(cls, _, fields, schema, environment):
        for field, constraints in fields.items():
            yield field, cls.solve_field(field, constraints, schema, environment)

    @classmethod
    def solve(cls, group, fields, schema, environment):
        solve_group = cls.solvers.get(group + ".", cls.solve_nogroup)
        for field, value in solve_group(group, fields, schema, environment):
            yield field, value


from . import boolean, date, geo_point, ip, keyword, long
