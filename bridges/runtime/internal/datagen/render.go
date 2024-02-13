// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"go.opentelemetry.io/otel/attribute"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

const (
	src  = "templates/*.tmpl"
	data = "data.go.tmpl"
)

func render(out string) (err error) {
	f, err := os.Create(out)
	if err != nil {
		return err
	}
	defer func() {
		if e := f.Close(); err == nil {
			err = e
		}
	}()

	var tmpl *template.Template
	funcMap := template.FuncMap{
		"attrStr": attrStr,
		"indent":  indent,
		"include": func(name string, data interface{}) (string, error) {
			buf := bytes.NewBuffer(nil)
			if err := tmpl.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
	}

	tmpl, err = template.New("base").Funcs(funcMap).ParseFS(templateFS, src)
	if err != nil {
		return err
	}
	err = tmpl.Lookup(data).Execute(f, Metrics)
	return err
}

func indent(tabs int, v string) string {
	pad := strings.Repeat("\t", tabs)
	return pad + strings.ReplaceAll(v, "\n", "\n"+pad)
}

func attrStr(a attribute.KeyValue) string {
	switch a.Value.Type() {
	case attribute.BOOL:
		return fmt.Sprintf("attribute.Bool(%q, %t)", a.Key, a.Value.AsBool())
	case attribute.INT64:
		return fmt.Sprintf("attribute.Int64(%q, %d)", a.Key, a.Value.AsInt64())
	case attribute.FLOAT64:
		return fmt.Sprintf("attribute.Float64(%q, %g)", a.Key, a.Value.AsFloat64())
	case attribute.STRING:
		return fmt.Sprintf("attribute.String(%q, %q)", a.Key, a.Value.AsString())
	case attribute.BOOLSLICE:
		return fmt.Sprintf("attribute.BoolSlice(%q, %#v)", a.Key, a.Value.AsBoolSlice())
	case attribute.INT64SLICE:
		return fmt.Sprintf("attribute.Int64Slice(%q, %#v)", a.Key, a.Value.AsInt64Slice())
	case attribute.FLOAT64SLICE:
		return fmt.Sprintf("attribute.Float64Slice(%q, %#v)", a.Key, a.Value.AsFloat64Slice())
	case attribute.STRINGSLICE:
		return fmt.Sprintf("attribute.StringSlice(%q, %#v)", a.Key, a.Value.AsStringSlice())
	}

	panic(fmt.Sprintf("invalid attribute type: %v", a.Value.Type()))
}
