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

// Condition represents node condition types
type Condition uint

// Clock represents a node clock type
type Clock uint

// Version represents a node version type
type Version uint

const (
	ConditionUnknown   = 0x00
	ConditionInitial   = 0x10
	ConditionBootstrap = 0x20
	ConditionReady     = 0x30
	ConditionStop      = 0x31
	ConditionOutOfDate = 0x32
)

// Status represents an abstract node interface for the status
type Status interface {
	// GetCondition returns the current status
	GetCondition() Condition
	// GetClock returns the current logical clock
	GetClock() Clock
	// GetVersion returns the current repository version
	GetVersion() Version
}
