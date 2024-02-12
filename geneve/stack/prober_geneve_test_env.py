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

"""Discover the stack used by geneve online tests"""

from ..utils.shelllib import ShellExpansionError
from .prober_elastic import ElasticStack


class GeneveTestEnvStack(ElasticStack):
    def __init__(self):
        config = {
            "name": "geneve-test-env",
            "elasticsearch": {
                "hosts": "$TEST_ELASTICSEARCH_URL",
                "api_key": "${TEST_API_KEY:-}",
                "ca_certs": "${TEST_CA_CERTS:-}",
                "verify_certs": "${TEST_VERIFY_CERTS:-}",
                "request_timeout": 30,
            },
            "kibana": {
                "url": "$TEST_KIBANA_URL",
                "api_key": "${TEST_API_KEY:-}",
                "ca_certs": "${TEST_CA_CERTS:-}",
                "verify_certs": "${TEST_VERIFY_CERTS:-}",
            },
        }
        super().__init__(config)


def probe():
    try:
        return [GeneveTestEnvStack()]
    except (ShellExpansionError, ValueError):
        return []


def load_from_config(config):
    pass
