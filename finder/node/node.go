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
	"crypto/sha256"
	"fmt"
)

// Node represents an abstract node interface
type Node interface {
	Config
	Status
	// UUID() returns a unique ID of the node
	UUID() string
}

// Equal returns true if the other node is same with this node
func Equal(this, other Node) bool {
	return ConfigEqual(this, other)
}

// GetUUID returns a unique ID with the specified node.
func GetUUID(node Node) string {
	seed := fmt.Sprintf("%s%s%d",
		node.Cluster(),
		node.Host(),
		node.RPCPort())
	h := sha256.New()
	h.Write([]byte(seed))
	return fmt.Sprintf("%x", h.Sum(nil))
}
