// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package geneve

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/geneve/cmd/control"
	"github.com/elastic/geneve/cmd/geneve/schema"
	"gopkg.in/yaml.v3"
)

var logger = log.New(log.Writer(), "datagen ", log.LstdFlags|log.Lmsgprefix)

type DocsSourceParams struct {
	Schema  string   `yaml:",omitempty"`
	Queries []string `yaml:",omitempty"`
}

type DocsSourceEntry struct {
	source DocsSource
	params DocsSourceParams
}

var dsEntriesMu = sync.Mutex{}
var dsEntries = make(map[string]DocsSourceEntry)

func getEntry(name string) (dse DocsSourceEntry, ok bool) {
	dsEntriesMu.Lock()
	defer dsEntriesMu.Unlock()
	dse, ok = dsEntries[name]
	return
}

func putEntry(name string, dse DocsSourceEntry) {
	dsEntriesMu.Lock()
	defer dsEntriesMu.Unlock()
	dsEntries[name] = dse
}

func delEntry(name string) {
	dsEntriesMu.Lock()
	defer dsEntriesMu.Unlock()
	delete(dsEntries, name)
}

func getSource(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var count int64 = 1
	var err error

	val := req.Form.Get("count")
	if val != "" {
		count, err = strconv.ParseInt(val, 10, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if count <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Count value must be greater than 0: %d\n", count)
			return
		}
	}

	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.Error(w, "Missing source name", http.StatusNotFound)
		return
	}

	name := parts[3]
	ds, ok := getEntry(name)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Documents source not found: %s\n", name)
		return
	}

	if len(parts) == 4 {
		enc := yaml.NewEncoder(w)
		if err := enc.Encode(ds.params); err != nil {
			http.Error(w, "Params encoding error", http.StatusInternalServerError)
			return
		}
		enc.Close()
		return
	}

	endpoint := parts[4]

	switch endpoint {
	case "_generate":
		docs, err := ds.source.Emit(int(count))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header()["Content-Type"] = []string{"application/json"}
		fmt.Fprintf(w, "[")
		for i, doc := range docs {
			if i > 0 {
				fmt.Fprintf(w, ",")
			}
			fmt.Fprintf(w, doc)
		}
		fmt.Fprintf(w, "]")
		return
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Unknown endpoint: %s\n", endpoint)
}

func getParamsFromRequest(w http.ResponseWriter, req *http.Request) (params DocsSourceParams, err error) {
	content_type, ok := req.Header["Content-Type"]
	if !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		err = fmt.Errorf("Missing Content-Type header")
		return
	}

	switch content_type[0] {
	case "application/yaml":
		err = yaml.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err.Error() == "EOF" {
				err = fmt.Errorf("No params were provided")
			}
		}

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
		err = fmt.Errorf("Unsupported Content-Type: %s", content_type[0])
	}

	return
}

func putSource(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.Error(w, "Missing source name", http.StatusNotFound)
		return
	}
	name := parts[3]

	params, err := getParamsFromRequest(w, req)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	var s schema.Schema
	if params.Schema != "" {
		var ok bool
		s, ok = schema.Get(params.Schema)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Schema not found: %s\n", params.Schema)
			return
		}
	}

	ds, err := NewDocsSource(s, params.Queries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	putEntry(name, DocsSourceEntry{source: ds, params: params})
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Created successfully")
	logger.Printf("%s %s", req.Method, req.URL)
}

func deleteSource(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.Error(w, "Missing source name", http.StatusNotFound)
		return
	}

	delEntry(parts[3])
	fmt.Fprintln(w, "Deleted successfully")
	logger.Printf("%s %s", req.Method, req.URL)
}

func getSchema(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.Error(w, "Missing schema name", http.StatusNotFound)
		return
	}

	name := parts[3]
	schema, ok := schema.Get(name)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Schema not found: %s\n", name)
		return
	}

	if len(parts) > 4 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unknown endpoint: %s\n", parts[4])
		return
	}

	enc := yaml.NewEncoder(w)
	if err := enc.Encode(schema); err != nil {
		http.Error(w, "Schema encoding error", http.StatusInternalServerError)
		return
	}
	enc.Close()
}

func getSchemaFromRequest(w http.ResponseWriter, req *http.Request) (schema schema.Schema, err error) {
	content_type, ok := req.Header["Content-Type"]
	if !ok {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		err = fmt.Errorf("Missing Content-Type header")
		return
	}

	switch content_type[0] {
	case "application/yaml":
		err = yaml.NewDecoder(req.Body).Decode(&schema)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err.Error() == "EOF" {
				err = fmt.Errorf("No schema was provided")
			}
		}

	default:
		w.WriteHeader(http.StatusUnsupportedMediaType)
		err = fmt.Errorf("Unsupported Content-Type: %s", content_type[0])
	}

	return
}

func putSchema(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.Error(w, "Missing schema name", http.StatusNotFound)
		return
	}
	name := parts[3]

	s, err := getSchemaFromRequest(w, req)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	schema.Put(name, s)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Created successfully")
	logger.Printf("%s %s", req.Method, req.URL)
}

func deleteSchema(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 4 || parts[3] == "" {
		http.Error(w, "Missing schema name", http.StatusNotFound)
		return
	}

	schema.Del(parts[3])
	fmt.Fprintln(w, "Deleted successfully")
	logger.Printf("%s %s", req.Method, req.URL)
}

func Use() {
	// this is just to bring this file in the compilation, the rest is done by init()
}

func init() {
	control.Handle("/api/docs_source/", &control.Handler{GET: getSource, PUT: putSource, DELETE: deleteSource})
	control.Handle("/api/schema/", &control.Handler{GET: getSchema, PUT: putSchema, DELETE: deleteSchema})
}
