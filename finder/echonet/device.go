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
	uecho_echonet "github.com/cybergarage/uecho-go/net/echonet"
	uecho_encoding "github.com/cybergarage/uecho-go/net/echonet/encoding"
	"github.com/cybergarage/uecho-go/net/echonet/log"
	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"
)

const (
	FinderDeviceCode = 0x0F9101
)

const (
	FinderConditionCode  = 0x80 // (Same as operation status code for Echonet)
	FinderClusterCode    = 0xA0
	FinderNameCode       = 0xA1
	FinderAddressCode    = 0xA2
	FinderRPCPortCode    = 0xA3
	FinderRenderPortCode = 0xA4
	FinderCarbonPortCode = 0xA5

	FinderClockCode   = 0xB0
	FinderVersionCode = 0xB1
)

const (
	FinderRPCPortSize    = 4
	FinderRenderPortSize = 4
	FinderCarbonPortSize = 4

	FinderClockSize   = 8
	FinderVersionSize = 8
)

func FinderDeviceAllPropertyCodes() []uecho_echonet.PropertyCode {
	props := []uecho_echonet.PropertyCode{
		FinderClusterCode,
		FinderNameCode,
		FinderAddressCode,
		FinderRPCPortCode,
		FinderRenderPortCode,
		FinderCarbonPortCode,
		// TODO : Send condition properties too if you neet
		//FinderConditionCode,
		//FinderClockCode,
		//FinderVersionCode,
	}
	return props
}

// EchonetDevice represents a base device for Echonet.
type EchonetDevice struct {
	*uecho_echonet.Device
}

// NewDevice returns a finder device.
func NewDevice() *EchonetDevice {
	dev := uecho_echonet.NewDevice()
	dev.SetCode(FinderDeviceCode)

	for _, propCode := range FinderDeviceAllPropertyCodes() {
		dev.CreateProperty(propCode, uecho_echonet.PropertyAttributeRead)
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
			propData = []byte{byte(node.GetCondition())}
		case FinderClusterCode:
			propData = []byte(node.GetCluster())
		case FinderNameCode:
			propData = []byte(node.GetName())
		case FinderAddressCode:
			propData = []byte(node.GetAddress())
		case FinderRPCPortCode:
			propData = make([]byte, FinderRPCPortSize)
			uecho_encoding.IntegerToByte(uint(node.GetRPCPort()), propData)
		case FinderRenderPortCode:
			propData = make([]byte, FinderRenderPortSize)
			uecho_encoding.IntegerToByte(uint(node.GetRenderPort()), propData)
		case FinderCarbonPortCode:
			propData = make([]byte, FinderCarbonPortSize)
			uecho_encoding.IntegerToByte(uint(node.GetCarbonPort()), propData)
		case FinderClockCode:
			propData = make([]byte, FinderClockSize)
			uecho_encoding.IntegerToByte(uint(node.GetClock()), propData)
		case FinderVersionCode:
			propData = make([]byte, FinderVersionSize)
			uecho_encoding.IntegerToByte(uint(node.GetVersion()), propData)
		default:
			continue
		}

		err := dev.SetPropertyData(uecho_protocol.PropertyCode(propCode), propData)
		if err != nil {
			log.Error(err.Error())
		}
	}
}
