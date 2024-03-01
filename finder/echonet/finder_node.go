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
	"net"

	"github.com/cybergarage/go-finder/finder/node"
	uecho "github.com/cybergarage/uecho-go/net/echonet"
)

const (
	errorEchonetFinderInvalidMessage         = "Invalid Echonet message : %s"
	errorEchonetFinderMessageInvalidObject   = "Invalid Echonet object code : %X != %X"
	errorEchonetFinderMessageInvalidProperty = "Invalid Echonet property code : %X"
)

type finderNode struct {
	*node.BaseNode
}

// NewFinderNodeWithResponseMesssage returns a new finder node with the specified message.
func NewFinderNodeWithResponseMesssage(msg *uecho.Message) (node.Node, error) {
	// Valdate the specified message

	if msg == nil {
		return nil, fmt.Errorf(errorEchonetFinderInvalidMessage, msg)
	}

	if msg.SEOJ() != FinderDeviceCode {
		return nil, fmt.Errorf(errorEchonetFinderMessageInvalidObject, msg.SEOJ(), FinderDeviceCode)
	}

	for _, propCode := range FinderDeviceAllPropertyCodes() {
		if !msg.HasProperty(propCode) {
			return nil, fmt.Errorf(errorEchonetFinderInvalidMessage, msg)
		}
	}

	// Create a candidate from the specified message

	candidateNode := &finderNode{
		BaseNode: node.NewBaseNode(),
	}

	for _, prop := range msg.Properties() {
		switch prop.Code() {
		case FinderConditionCode:
			candidateNode.SetCondition(node.Condition(prop.IntegerData()))
		case FinderClusterCode:
			candidateNode.SetCluster(prop.StringData())
		case FinderHostCode:
			candidateNode.SetHost(prop.StringData())
		case FinderAddressCode:
			candidateNode.SetAddress(net.ParseIP(prop.StringData()))
		case FinderRPCPortCode:
			candidateNode.SetRPCPort(prop.IntegerData())
		case FinderClockCode:
			candidateNode.SetClock(node.Clock(prop.IntegerData()))
		default:
			return nil, fmt.Errorf(errorEchonetFinderMessageInvalidProperty, prop.Code())
		}
	}
	return candidateNode, nil
}
