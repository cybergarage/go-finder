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
	"fmt"
	"regexp"
	"testing"

	"github.com/cybergarage/go-finder/finder/node"
)

const (
	testFinderNodeCountError     = "Node count error : %d != %d"
	testFinderMatchingError      = "Matching error (%s) : %s"
	testFinderMatchingCountError = "Matching count error (%s) : %d != %d"
)

var testFinderNodeNames = []string{
	"org.cybergarage.finder001",
	"org.cybergarage.finder002",
	"org.cybergarage.finder003",
}

func setupTestFinderNodes() []Node {
	nodes := make([]Node, len(testFinderNodeNames))
	for n, name := range testFinderNodeNames {
		node := node.NewBaseNode()
		node.SetHost(name)
		nodes[n] = node
	}
	return nodes
}

func finderTest(t *testing.T, finder Finder) error {
	// Check all nodes

	nodes, err := finder.GetAllNodes()
	if err != nil {
		return err
	}

	if len(nodes) != len(testFinderNodeNames) {
		return fmt.Errorf(testFinderNodeCountError, len(nodes), len(testFinderNodeNames))
	}

	// Check regexp names for a node

	for _, nodeName := range testFinderNodeNames {
		var err error
		var nodes []node.Node
		t.Run(nodeName, func(t *testing.T) {
			re := regexp.MustCompile(nodeName)
			nodes, err = finder.GetRegexpNodes(re)
			if err != nil {
				return
			}
			if len(nodes) != 1 {
				err = fmt.Errorf(testFinderNodeCountError, len(nodes), 1)
				return
			}
			if nodes[0].Host() != nodeName {
				err = fmt.Errorf(testFinderMatchingError, nodeName, nodes[0].Host())
				return
			}
		})
		if err != nil {
			return err
		}
	}

	// Check regexp names for all nodes

	reNames := []string{
		".*",
		"^org.cybergarage.finder",
		"org.cybergarage.finder.*",
		"org.cybergarage.finder00[1-3]",
	}

	for _, reName := range reNames {
		var err error
		var nodes []node.Node
		t.Run(reName, func(t *testing.T) {
			re := regexp.MustCompile(reName)
			nodes, err = finder.GetRegexpNodes(re)
			if err != nil {
				return
			}
			if len(nodes) != len(testFinderNodeNames) {
				err = fmt.Errorf(testFinderMatchingCountError, reName, len(nodes), len(testFinderNodeNames))
				return
			}
		})
		if err != nil {
			return err
		}
	}

	// Check start names for all nodes

	metricsNames := []string{
		"org.cybergarage.finder001",
		"org.cybergarage.finder001.m1",
		"org.cybergarage.finder001.system.m1",
		"org.cybergarage.finder002",
		"org.cybergarage.finder002.m1",
		"org.cybergarage.finder002.system.m1",
		"org.cybergarage.finder003",
		"org.cybergarage.finder003.m1",
		"org.cybergarage.finder003.system.m1",
	}

	for _, metricsName := range metricsNames {
		var err error
		var nodes []node.Node
		t.Run(metricsName, func(t *testing.T) {
			nodes, err = finder.GetPrefixNodes(metricsName)
			if err != nil {
				return
			}
			if len(nodes) != 1 {
				err = fmt.Errorf(testFinderMatchingCountError, metricsName, len(nodes), 1)
				return
			}
		})
		if err != nil {
			return err
		}
	}

	return nil
}
