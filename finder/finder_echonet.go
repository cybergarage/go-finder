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
	"reflect"
	"time"

	uecho "github.com/cybergarage/uecho-go/net/echonet"
	uecho_protocol "github.com/cybergarage/uecho-go/net/echonet/protocol"

	finder_echonet "github.com/cybergarage/go-finder/finder/echonet"
	"github.com/cybergarage/go-finder/finder/node"
	"github.com/cybergarage/go-logger/log"
)

const (
	echonetFinderSearchSleepSecond = 1
)

const (
	errorEchonetFinderNoResponse     = "Echonet node (%s:%d) is not responding"
	msgEchonetFinderFoundEchonetNode = "Echonet node (%s:%d) is found"
	msgEchonetFinderFoundCadiateNode = "Candidate finder node (%s:%d) is found"
	msgEchonetFinderFoundNewNode     = "New finder node (%s:%d) is found"
)

// EchonetFinder represents a base finder.
type EchonetFinder struct {
	*baseFinder
	localNode node.Node
	*finder_echonet.EchonetController
}

// NewEchonetFinderWithLocalNode returns a new finder with the specified node.
func NewEchonetFinderWithLocalNode(node node.Node) Finder {
	finder := &EchonetFinder{
		baseFinder:        newBaseFinder(),
		localNode:         node,
		EchonetController: finder_echonet.NewController(),
	}
	finder.EchonetController.SetListener(finder)
	return finder
}

// NewEchonetFinder returns a new finder of Echonet.
func NewEchonetFinder() Finder {
	return NewEchonetFinderWithLocalNode(nil)
}

// Search searches all nodes.
func (finder *EchonetFinder) Search() error {
	err := finder.EchonetController.SearchAllObjects()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * echonetFinderSearchSleepSecond)
	return nil
}

// IsLocalNode returns true when the specified node is the local node, otherwise false.
func (finder *EchonetFinder) IsLocalNode(candidateNode node.Node) bool {
	if finder.localNode == nil || reflect.ValueOf(finder.localNode).IsNil() {
		return false
	}

	return node.Equal(finder.localNode, candidateNode)
}

// Start starts the finder.
func (finder *EchonetFinder) Start() error {
	return finder.EchonetController.Start()
}

// Stop stops the finder.
func (finder *EchonetFinder) Stop() error {
	return finder.EchonetController.Stop()
}

// IsRunning returns true when the finder is running, otherwise false.
func (finder *EchonetFinder) IsRunning() bool {
	return finder.EchonetController.IsRunning()
}

// String returns the description
func (finder *EchonetFinder) String() string {
	return fmt.Sprintf("%s:%s", FinderEchonet, uecho.Version)
}

func (finder *EchonetFinder) ControllerMessageReceived(msg *uecho_protocol.Message) {
	if !msg.IsReadRequest() {
		return
	}

	finder.EchonetController.EchonetDevice.UpdatePropertyWithNode(finder.localNode)
}

func (finder *EchonetFinder) ControllerNewNodeFound(echonetNode *uecho.RemoteNode) {
	if !finder.IsRunning() {
		return
	}

	log.Trace(msgEchonetFinderFoundEchonetNode, echonetNode.Address(), echonetNode.Port())

	reqMsg := finder_echonet.NewRequestAllPropertiesMessage()
	resMsg, err := finder.EchonetController.PostMessage(echonetNode, reqMsg)
	if err != nil {
		log.Error(errorEchonetFinderNoResponse, echonetNode.Address(), echonetNode.Port())
		log.Error("%s", err.Error())
		return
	}

	candidateNode, err := finder_echonet.NewFinderNodeWithResponseMesssage(resMsg)
	if err != nil {
		log.Error("%s", err.Error())
		return
	}

	log.Trace(msgEchonetFinderFoundCadiateNode, candidateNode.Address(), candidateNode.RPCPort())

	if finder.IsLocalNode(candidateNode) {
		return
	}

	if finder.HasNode(candidateNode) {
		return
	}

	log.Info(msgEchonetFinderFoundNewNode, candidateNode.Address(), candidateNode.RPCPort())

	err = finder.addNode(candidateNode)
	if err != nil {
		log.Error("%s", err.Error())
		return
	}
}
