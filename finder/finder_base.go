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
	"math/rand"
	"regexp"
	"strings"

	"github.com/cybergarage/go-finder/finder/node"
)

const (
	errorFinderHasNoNodes  = "Finder hasnt' find any nodes"
	errorFinderHasSameNode = "Node (%s) is already added"
)

// baseFinder represents a base finder.
type baseFinder struct {
	nodes          []Node
	searchListener FinderSearchListener
	notifyListener FinderNotifyListener
}

// newBaseFinder returns a new base finder.
func newBaseFinder() *baseFinder {
	finder := &baseFinder{
		nodes:          make([]Node, 0),
		searchListener: nil,
		notifyListener: nil,
	}
	return finder
}

// SetSearchListener sets the search listener.
func (finder *baseFinder) SetSearchListener(l FinderSearchListener) error {
	finder.searchListener = l
	return nil
}

// SetSearchListener sets the search listener.
func (finder *baseFinder) SetNotifyListener(l FinderNotifyListener) error {
	finder.notifyListener = l
	return nil
}

// HasNode returns true when the specified node is added already, otherwise false.
func (finder *baseFinder) HasNode(targetNode Node) bool {
	for _, addedNode := range finder.nodes {
		if node.Equal(targetNode, addedNode) {
			return true
		}
	}
	return false
}

// addNodes adds a specified node.
func (finder *baseFinder) addNode(node Node) error {
	if finder.HasNode(node) {
		return fmt.Errorf(errorFinderHasSameNode, node)
	}
	finder.nodes = append(finder.nodes, node)
	return nil
}

// GetAllNodes returns all found nodes.
func (finder *baseFinder) GetAllNodes() ([]Node, error) {
	return finder.nodes, nil
}

// GetNeighborhoodNode returns a neighborhood node of the specified node.
func (finder *baseFinder) GetNeighborhoodNode(node Node) (Node, error) {
	nodes, err := finder.GetAllNodes()
	if err != nil {
		return nil, err
	}
	nodeCount := len(nodes)
	if nodeCount <= 0 {
		return nil, fmt.Errorf(errorFinderHasNoNodes)
	}

	// FIXME : Return  a neighborhood node of the specified node instead of the random node
	nodeIdx := rand.Int() % nodeCount

	return nodes[nodeIdx], nil
}

// GetPrefixNodes returns only nodes matching with a specified start string
func (finder *baseFinder) GetPrefixNodes(targetString string) ([]Node, error) {
	nodes, err := finder.GetAllNodes()
	if err != nil {
		return nil, err
	}

	matchedNodes := make([]Node, 0)

	for _, node := range nodes {
		port := node.RPCPort()
		addr := node.Address()
		name := node.Host()

		hosts := []string{
			addr.String(),
			fmt.Sprintf("%s:%d", addr, port),
			name,
			fmt.Sprintf("%s:%d", name, port),
		}

		for _, host := range hosts {
			if len(host) <= 0 {
				continue
			}
			if strings.HasPrefix(targetString, host) {
				matchedNodes = append(matchedNodes, node)
				break
			}
		}
	}

	return matchedNodes, nil
}

// GetRegexpNodes returns only nodes matching with a specified regular expression
func (finder *baseFinder) GetRegexpNodes(re *regexp.Regexp) ([]Node, error) {
	nodes, err := finder.GetAllNodes()
	if err != nil {
		return nil, err
	}

	matchedNodes := make([]Node, 0)

	for _, node := range nodes {
		port := node.RPCPort()
		addr := node.Address()
		name := node.Host()

		hosts := []string{
			addr.String(),
			fmt.Sprintf("%s:%d", addr, port),
			name,
			fmt.Sprintf("%s:%d", name, port),
		}

		for _, host := range hosts {
			if len(host) <= 0 {
				continue
			}
			if re.MatchString(host) {
				matchedNodes = append(matchedNodes, node)
				break
			}
		}
	}

	return matchedNodes, nil
}
