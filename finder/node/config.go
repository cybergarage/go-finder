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

// Config represents an abstract node interface for the configuration
type Config interface {
	// GetCluster returns the cluster name
	GetCluster() string
	// GetName returns the host name
	GetName() string
	// GetAddress returns the interface address
	GetAddress() string
	// GetRPCPort returns the RPC port
	GetRPCPort() int
	// GetRenderPort returns the Graphite render port
	GetRenderPort() int
	// GetCarbonPort returns the Graphite carbon port
	GetCarbonPort() int
}

// ConfigEqual returns true if the other node is same with this node
func ConfigEqual(this, other Config) bool {
	if this.GetCluster() != other.GetCluster() {
		return false
	}

	if 0 < len(this.GetName()) && 0 < len(other.GetName()) {
		if this.GetName() != other.GetName() {
			return false
		}
	}

	if 0 < len(this.GetAddress()) && 0 < len(other.GetAddress()) {
		if this.GetAddress() != other.GetAddress() {
			return false
		}
	}

	if this.GetRPCPort() != other.GetRPCPort() {
		return false
	}

	if this.GetRenderPort() != other.GetRenderPort() {
		return false
	}

	if this.GetCarbonPort() != other.GetCarbonPort() {
		return false
	}

	return true
}
