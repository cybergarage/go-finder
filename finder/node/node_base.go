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
)

// BaseNode represents a base node.
type BaseNode struct {
	Node
	Cluster    string
	Name       string
	Address    string
	RPCPort    int
	RenderPort int
	CarbonPort int
	Clock      Clock
	Condition  Condition
	Version    Version
}

// NewBaseNode returns a new base node.
func NewBaseNode() *BaseNode {
	node := &BaseNode{
		Condition: ConditionInitial,
		Clock:     0,
		Version:   0,
	}
	return node
}

// NewNode returns a new node.
func NewNode() Node {
	return NewBaseNode()
}

// UpdateClock increments the internal clock
func (node *BaseNode) UpdateClock() {
	node.Clock++
}

// UpdateVersion increments the internal version
func (node *BaseNode) UpdateVersion() {
	node.Version++
}

// SetStatus sets a new status
func (node *BaseNode) SetStatus(status Status) {
	node.Clock = status.GetClock()
	node.Condition = status.GetCondition()
	node.Version = status.GetVersion()
}

// GetCluster returns the cluster name
func (node *BaseNode) GetCluster() string {
	return node.Cluster
}

// GetName returns the host name
func (node *BaseNode) GetName() string {
	if 0 < len(node.Name) {
		return node.Name
	}

	if len(node.Address) <= 0 {
		return ""
	}
	names, err := net.LookupAddr(node.Address)
	if err != nil {
		return ""
	}
	node.Name = names[0]

	return node.Name
}

// GetAddress returns the interface address
func (node *BaseNode) GetAddress() string {
	if 0 < len(node.Address) {
		return node.Address
	}

	if len(node.Name) <= 0 {
		return ""
	}
	addrs, err := net.LookupIP(node.Name)
	if err != nil {
		return ""
	}
	node.Address = addrs[0].String()

	return node.Address
}

// GetRPCPort returns the RPC port
func (node *BaseNode) GetRPCPort() int {
	return node.RPCPort
}

// GetRenderPort returns the Graphite render port
func (node *BaseNode) GetRenderPort() int {
	return node.RenderPort
}

// GetCarbonPort returns the Graphite carbon port
func (node *BaseNode) GetCarbonPort() int {
	return node.CarbonPort
}

// GetCondition returns the current status
func (node *BaseNode) GetCondition() Condition {
	return node.Condition
}

// GetClock returns the current logical clock
func (node *BaseNode) GetClock() Clock {
	return node.Clock
}

// GetVersion returns the current repository version
func (node *BaseNode) GetVersion() Version {
	return node.Version
}

// GetUniqueID returns a unique ID of the node
func (node *BaseNode) GetUniqueID() string {
	return GetUniqueID(node)
}
