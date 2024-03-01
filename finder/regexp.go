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
	"regexp"
	"strings"
)

const (
	// See : http://graphite.readthedocs.io/en/latest/render_api.html
	exprAsteriskGraphite = "*"
	exprAsteriskGo       = "(.*)"
)

// Regexp represents a regexp for the finder.
type Regexp struct {
	expr     string
	goRegexp *regexp.Regexp
}

// NewRegexp returns a new regexp.
func NewRegexp() *Regexp {
	regexp := &Regexp{
		goRegexp: nil,
	}
	return regexp
}

// Compile parses a regular expression.
func (re *Regexp) Compile(expr string) error {
	var err error
	re.goRegexp, err = regexp.Compile(expr)
	if err != nil {
		return err
	}
	re.expr = expr
	return nil
}

// CompileGraphite parses a regular expression to Graphite
// See : http://graphite.readthedocs.io/en/latest/render_api.html
func (re *Regexp) CompileGraphite(expr string) error {
	// FIXME : Only replacing the prefix expression string
	if strings.HasPrefix(expr, exprAsteriskGraphite) {
		expr = strings.Replace(expr, exprAsteriskGraphite, exprAsteriskGo, 1)
	}

	return re.Compile(expr)
}

// matchNodeString reports whether the Regexp matches the string.
func (re *Regexp) matchNodeString(nodeStr string) bool {
	if len(nodeStr) <= 0 {
		return false
	}

	// FIXME : Only checking the prefix expression string
	if strings.HasPrefix(re.expr, exprAsteriskGo) {
		return true
	}
	if strings.HasPrefix(re.expr, nodeStr) {
		return true
	}

	return re.goRegexp.MatchString(nodeStr)
}

// MatchNode reports whether the Regexp matches the node.
func (re *Regexp) MatchNode(node Node) bool {
	ok := re.matchNodeString(node.Host())
	if ok {
		return true
	}

	return re.matchNodeString(node.Address().String())
}

// expandNodeString replaces expression to node returns the result;.
func (re *Regexp) expandNodeString(nodeStr string) (string, bool) {
	if len(nodeStr) <= 0 {
		return "", false
	}

	// FIXME : Only replacing the prefix expression string
	if strings.HasPrefix(re.expr, exprAsteriskGo) {
		return strings.Replace(re.expr, exprAsteriskGo, nodeStr, 1), true
	}

	return re.expr, true
}

// ExpandNode replaces expression to node returns the result;.
func (re *Regexp) ExpandNode(node Node) (string, bool) {
	result, ok := re.expandNodeString(node.Host())
	if ok {
		return result, true
	}

	result, ok = re.expandNodeString(node.Address().String())
	if ok {
		return result, true
	}

	return "", false
}
