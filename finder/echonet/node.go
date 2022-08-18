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

package echonet

import (
	"fmt"
	"reflect"

	"github.com/cybergarage/go-finder/finder/node"
	uecho_echonet "github.com/cybergarage/uecho-go/net/echonet"
	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	errorNodeNotRunning = "Node is not running"
)

type EchonetNode struct {
	*uecho_echonet.LocalNode
	*EchonetDevice
	node.Node
}

// NewEchonetNodeWithNode returns a new finder node.
func NewEchonetNodeWithNode(srcNode node.Node) (*EchonetNode, error) {
	node := &EchonetNode{
		LocalNode:     uecho_echonet.NewLocalNode(),
		EchonetDevice: NewDevice(),
		Node:          srcNode,
	}

	node.SetConfig(NewDefaultConfig())
	node.SetManufacturerCode(ManufacturerCode)

	node.SetListener(node)
	node.AddDevice(node.EchonetDevice.Device)

	return node, nil
}

// GetAddress returns the interface address
func (node *EchonetNode) GetAddress() string {
	return node.LocalNode.GetAddress()
}

// GetLocalNode returns the local echonet node in the node
func (node *EchonetNode) GetLocalNode() *uecho_echonet.LocalNode {
	return node.LocalNode
}

// GetLocalDevice returns the local echonet device in the node.
func (node *EchonetNode) GetLocalDevice() *uecho_echonet.Device {
	return node.EchonetDevice.Device
}

// HasSourceNode returns true when this node has a source node, otherwise false.
func (node *EchonetNode) HasSourceNode() bool {
	if node.Node == nil || reflect.ValueOf(node.Node).IsNil() {
		return false
	}
	return true
}

// GetSourceNode returns the local souce node in the node.
func (node *EchonetNode) GetSourceNode() node.Node {
	return node.Node
}

// NodeMessageReceived updates local properties for the source node.
func (node *EchonetNode) NodeMessageReceived(msg *uecho_protocol.Message) error {
	if !node.IsRunning() {
		return fmt.Errorf(errorNodeNotRunning)
	}

	if !msg.IsReadRequest() {
		return nil
	}

	if !node.HasSourceNode() {
		return nil
	}

	node.EchonetDevice.UpdatePropertyWithNode(node)

	return nil
}
