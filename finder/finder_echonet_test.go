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
	"time"

	"github.com/cybergarage/go-finder/finder/echonet"
	"github.com/cybergarage/go-logger/log"
)

func setupTestEchonetFinderNodes() ([]*echonet.EchonetNode, error) {
	nodes := setupTestFinderNodes()
	echonetNodes := make([]*echonet.EchonetNode, len(nodes))
	for n, node := range nodes {
		echonetNode, err := echonet.NewEchonetNodeWithNode(node)
		if err != nil {
			return nil, err
		}
		echonetNodes[n] = echonetNode
	}
	return echonetNodes, nil
}

func finderEchonetTest(t *testing.T, finder Finder, nodes []*echonet.EchonetNode) {
	// Check all nodes

	_, err := finder.GetAllNodes()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestEchonetFinder(t *testing.T) {
	log.SetStdoutDebugEnbled(true)
	defer log.SetStdoutDebugEnbled(false)

	nodes, err := setupTestEchonetFinderNodes()
	if err != nil {
		t.Error(err)
		return
	}

	for n, node := range nodes {
		err = node.Start()
		if err != nil {
			t.Error(err)
		}
		log.Tracef("node[%d]: %s:%d", n, node.Address(), node.Port())
	}

	finder := NewEchonetFinder()

	err = finder.Start()
	if err != nil {
		t.Error(err)
		return
	}

	err = finder.Search()
	if err != nil {
		t.Error(err)
		return
	}

	time.Sleep((500 * time.Millisecond) * time.Duration(len(nodes)))

	// FIXME : Update uecho-go to be able to the neighborhood node on CentOS
	err = finderTest(t, finder)
	if err != nil {
		t.Skip(err)
	}
	finderEchonetTest(t, finder, nodes)

	err = finder.Stop()
	if err != nil {
		t.Error(err)
	}

	for _, node := range nodes {
		err = node.Stop()
		if err != nil {
			t.Error(err)
			return
		}
	}
}
