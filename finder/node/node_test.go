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

package node

import (
	"regexp"
	"testing"
)

const (
	testNodeName          = "finder001.cybergarage.org"
	testNodeMatchingError = "Matching Error (%s) : %s"
)

var testNodeRegExpStrings = []string{
	testNodeName,
	".*",
}

func nodeMachingTest(t *testing.T, node Node) {
	name := node.GetName()
	for _, reString := range testNodeRegExpStrings {
		re := regexp.MustCompile(reString)
		if !re.MatchString(name) {
			t.Errorf(testNodeMatchingError, reString, name)
		}
	}
}

func TestNewBaseNode(t *testing.T) {
	node := NewBaseNode()
	node.Name = testNodeName
	nodeMachingTest(t, node)
}

func TestNodeHasName(t *testing.T) {
	node := NewBaseNode()
	node.Address = "127.0.0.1"
	if len(node.GetName()) <= 0 {
		t.Errorf("No name : %s", node.GetAddress())
	}
}

func TestNodeHasAddress(t *testing.T) {
	node := NewBaseNode()
	node.Name = "localhost"
	if len(node.GetAddress()) <= 0 {
		t.Errorf("No address : %s", node.GetName())
	}
}

func TestEqual(t *testing.T) {
	node01 := NewBaseNode()
	node02 := NewBaseNode()

	// name

	node01.Name = "node01"
	node01.Address = ""
	node02.Name = "node02"
	node02.Address = ""

	if !Equal(node01, node01) {
		t.Errorf("%s != %s", node01.GetName(), node01.GetName())
	}

	if Equal(node01, node02) {
		t.Errorf("%s == %s", node01.GetName(), node02.GetName())
	}

	// address

	node01.Name = ""
	node01.Address = "192.168.100.1"
	node02.Name = ""
	node02.Address = "192.168.100.2"

	if !Equal(node01, node01) {
		t.Errorf("%s != %s", node01.GetName(), node01.GetName())
	}

	if Equal(node01, node02) {
		t.Errorf("%s == %s", node01.GetName(), node02.GetName())
	}

	// port

	node01.Name = "node01"
	node01.Address = "192.168.100.1"
	node01.RPCPort = 0001
	node02.Name = node01.Name
	node02.Address = node01.Address
	node02.RPCPort = 0002

	if !Equal(node01, node01) {
		t.Errorf("%s != %s", node01.GetName(), node01.GetName())
	}

	if Equal(node01, node02) {
		t.Errorf("%s == %s", node01.GetName(), node02.GetName())
	}
}
