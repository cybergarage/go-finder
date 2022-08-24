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
	"net"
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
	name := node.Host()
	for _, reString := range testNodeRegExpStrings {
		re := regexp.MustCompile(reString)
		if !re.MatchString(name) {
			t.Errorf(testNodeMatchingError, reString, name)
		}
	}
}

func TestNewBaseNode(t *testing.T) {
	node := NewBaseNode().SetHost(testNodeName)
	nodeMachingTest(t, node)
}

func TestNodeHasName(t *testing.T) {
	node := NewBaseNode().SetAddress(net.ParseIP("127.0.0.1"))
	if len(node.Host()) <= 0 {
		t.Errorf("No host : %s", node.Address())
	}
}

func TestNodeHasAddress(t *testing.T) {
	node := NewBaseNode().SetHost("localhost")
	if node.Address() == nil {
		t.Errorf("No address : %s", node.Host())
	}
}

func TestEqual(t *testing.T) {
	node01 := NewBaseNode().SetHost("node01").SetAddress(nil)
	node02 := NewBaseNode().SetHost("node02").SetAddress(nil)

	// name

	if !Equal(node01, node01) {
		t.Errorf("%s != %s", node01.Host(), node01.Host())
	}

	if Equal(node01, node02) {
		t.Errorf("%s == %s", node01.Host(), node02.Host())
	}

	// address

	node01 = NewBaseNode().SetHost("").SetAddress(net.ParseIP("192.168.100.1"))
	node02 = NewBaseNode().SetHost("").SetAddress(net.ParseIP("192.168.100.2"))

	if !Equal(node01, node01) {
		t.Errorf("%s != %s", node01.Host(), node01.Host())
	}

	if Equal(node01, node02) {
		t.Errorf("%s == %s", node01.Host(), node02.Host())
	}

	// port

	node01 = NewBaseNode().SetHost("node01").SetAddress(net.ParseIP("192.168.100.1")).SetRPCPort(0001)
	node02 = NewBaseNode().SetHost("node01").SetAddress(net.ParseIP("192.168.100.1")).SetRPCPort(0002)

	if !Equal(node01, node01) {
		t.Errorf("%s != %s", node01.Host(), node01.Host())
	}

	if Equal(node01, node02) {
		t.Errorf("%s == %s", node01.Host(), node02.Host())
	}
}
