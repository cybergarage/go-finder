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
	"reflect"

	"github.com/cybergarage/go-finder/finder/node"
	"github.com/cybergarage/go-logger/log"
	uecho "github.com/cybergarage/uecho-go/net/echonet"
	uecho_encoding "github.com/cybergarage/uecho-go/net/echonet/encoding"
	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	FinderDeviceCode = 0x0F9101
)

const (
	FinderConditionCode = 0x80 // (Same as operation status code for Echonet)
	FinderClusterCode   = 0xA0
	FinderHostCode      = 0xA1
	FinderAddressCode   = 0xA2
	FinderRPCPortCode   = 0xA3
	FinderClockCode     = 0xB0
)

const (
	FinderRPCPortSize    = 4
	FinderRenderPortSize = 4
	FinderCarbonPortSize = 4

	FinderClockSize   = 8
	FinderVersionSize = 8
)

func FinderDeviceAllPropertyCodes() []uecho.PropertyCode {
	props := []uecho.PropertyCode{
		FinderClusterCode,
		FinderHostCode,
		FinderAddressCode,
		FinderRPCPortCode,
		FinderClockCode,
	}
	return props
}

// EchonetDevice represents a base device for Echonet.
type EchonetDevice struct {
	*uecho.Device
}

// NewDevice returns a finder device.
func NewDevice() *EchonetDevice {
	dev := uecho.NewDevice()
	dev.SetCode(FinderDeviceCode)

	for _, propCode := range FinderDeviceAllPropertyCodes() {
		dev.AddProperty(uecho.NewPropertyWithCode(propCode).SetReadAttribute(uecho.Required))
	}

	return &EchonetDevice{Device: dev}
}

// UpdatePropertyWithNode updates the device property with the specified node.
func (dev *EchonetDevice) UpdatePropertyWithNode(node node.Node) {
	if node == nil || reflect.ValueOf(node).IsNil() {
		return
	}

	for _, propCode := range FinderDeviceAllPropertyCodes() {
		var propData []byte
		switch propCode {
		case FinderConditionCode:
			propData = []byte{byte(node.Condition())}
		case FinderClusterCode:
			propData = []byte(node.Cluster())
		case FinderHostCode:
			propData = []byte(node.Host())
		case FinderAddressCode:
			propData = []byte(node.Address())
		case FinderRPCPortCode:
			propData = make([]byte, FinderRPCPortSize)
			uecho_encoding.IntegerToByte(uint(node.RPCPort()), propData)
		case FinderClockCode:
			propData = make([]byte, FinderClockSize)
			uecho_encoding.IntegerToByte(uint(node.Clock()), propData)
		default:
			continue
		}

		err := dev.SetPropertyData(uecho_protocol.PropertyCode(propCode), propData)
		if err != nil {
			log.Error(err.Error())
		}
	}
}
