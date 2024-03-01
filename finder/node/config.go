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

import "net"

// Config represents an abstract node interface for the configuration.
type Config interface {
	// Cluster returns the cluster name.
	Cluster() string
	// Host returns the host name.
	Host() string
	// Host returns the address.
	Address() net.IP
	// RPCPort returns the RPC port.
	RPCPort() uint
}

// ConfigEqual returns true if the other node is same with this node.
func ConfigEqual(this, other Config) bool {
	if this.Cluster() != other.Cluster() {
		return false
	}

	if this.Host() != other.Host() {
		return false
	}

	if this.Address() == nil {
		if other.Address() != nil {
			return false
		}
	}

	if !this.Address().Equal(other.Address()) {
		return false
	}

	if this.RPCPort() != other.RPCPort() {
		return false
	}

	return true
}
