// Copyright (C) 2022 Satoshi Konno All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package finder

import (
	"testing"

	"github.com/cybergarage/go-finder/finder/node"
)

var testRegexTestCases = [][]string{
	// Graphite Wildcards (Only prefix wildcard)
	// See : http://graphite.readthedocs.io/en/latest/render_api.html
	{"node01", "*", "node01"},
	{"node01", "*.metrics01", "node01.metrics01"},
	{"node01", "*.service.metrics01", "node01.service.metrics01"},
	// Graphite Wildcards (With metrics wildcards)
	{"node01", "*.*", "node01.*"},
	{"node01", "*.metrics01.*", "node01.metrics01.*"},
	{"node01", "*.service.metrics01.*", "node01.service.metrics01.*"},
	{"node01", "*.service.*.metrics01.*", "node01.service.*.metrics01.*"},
	// Graphite Wildcards (No prefix wildcard)
	{"node01", "node01.*", "node01.*"},
	{"node01", "node01.metrics01.*", "node01.metrics01.*"},
	{"node01", "node01.service.metrics01.*", "node01.service.metrics01.*"},
	{"node01", "node01.service.*.metrics01.*", "node01.service.*.metrics01.*"},
}

func TestRegexpGraphite(t *testing.T) {
	for n, testCase := range testRegexTestCases {
		name := testCase[0]
		expr := testCase[1]
		expand := testCase[2]

		re := NewRegexp()
		err := re.CompileGraphite(expr)
		if err != nil {
			t.Errorf("[%d] %s : %s", n, expr, err)
			continue
		}

		node := node.NewBaseNode()
		node.Name = name

		ok := re.MatchNode(node)
		if !ok {
			t.Errorf("[%d] %s != %s", n, expr, node.GetName())
			continue
		}

		expandedName, ok := re.ExpandNode(node)
		if !ok || expand != expandedName {
			t.Errorf("[%d] %s != %s", n, expand, expandedName)
			continue
		}
	}
}
