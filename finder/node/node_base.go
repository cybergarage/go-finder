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
	cluster string
	host    string
	address net.IP
	rpcPort uint
	clock   Clock
	cond    Condition
}

// NewBaseNode returns a new base node.
func NewBaseNode() *BaseNode {
	node := &BaseNode{
		cond:  ConditionInitial,
		clock: 0,
	}
	return node
}

// NewNode returns a new node.
func NewNode() Node {
	return NewBaseNode()
}

// UpdateClock increments the internal clock.
func (node *BaseNode) UpdateClock() {
	node.clock++
}

// SetStatus sets the specified status to the node.
func (node *BaseNode) SetStatus(status Status) {
	node.clock = status.Clock()
	node.cond = status.Condition()
}

// SetCluster sets the specified cluster name to the node.
func (node *BaseNode) SetCluster(name string) *BaseNode {
	node.cluster = name
	return node
}

// SetHost sets the specified host name to the node.
func (node *BaseNode) SetHost(name string) *BaseNode {
	node.host = name
	return node
}

// SetAddress sets the specified address name to the node.
func (node *BaseNode) SetAddress(addr net.IP) *BaseNode {
	node.address = addr
	return node
}

// SetRPCPort sets the specified portto the node.
func (node *BaseNode) SetRPCPort(port uint) *BaseNode {
	node.rpcPort = port
	return node
}

// SetClock sets the specified clock to the node.
func (node *BaseNode) SetClock(val Clock) {
	node.clock = val
}

// SetCondition sets the specified condition to the node.
func (node *BaseNode) SetCondition(val Condition) {
	node.cond = val
}

// Cluster returns the cluster name.
func (node *BaseNode) Cluster() string {
	return node.cluster
}

// Host returns the host name.
func (node *BaseNode) Host() string {
	if 0 < len(node.host) {
		return node.host
	}

	if len(node.address) <= 0 {
		return ""
	}
	names, err := net.LookupAddr(node.address.String())
	if err != nil {
		return ""
	}
	node.host = names[0]

	return node.host
}

// Address returns the interface address.
func (node *BaseNode) Address() net.IP {
	if 0 < len(node.address) {
		return node.address
	}

	if len(node.host) <= 0 {
		return nil
	}
	addrs, err := net.LookupIP(node.host)
	if err != nil {
		return nil
	}
	node.address = addrs[0]

	return node.address
}

// RPCPort returns the RPC port.
func (node *BaseNode) RPCPort() uint {
	return node.rpcPort
}

// Condition returns the current status.
func (node *BaseNode) Cndition() Condition {
	return node.cond
}

// Clock returns the current logical clock.
func (node *BaseNode) Clock() Clock {
	return node.clock
}

// UUID returns a unique ID of the node.
func (node *BaseNode) UUID() string {
	return GetUUID(node)
}
