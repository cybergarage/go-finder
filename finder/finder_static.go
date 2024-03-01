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
	"github.com/cybergarage/go-logger/log"
)

// StaticFinder represents a simple static finder.
type StaticFinder struct {
	*baseFinder
}

// NewStaticFinderWithNodes returns a new static finder with specified nodes.
func NewStaticFinderWithNodes(nodes []Node) Finder {
	finder := &StaticFinder{
		baseFinder: newBaseFinder(),
	}

	for _, node := range nodes {
		err := finder.addNode(node)
		if err != nil {
			log.Errorf(err.Error())
		}
	}

	return finder
}

// SearchAll searches all nodes.
func (finder *StaticFinder) Search() error {
	return nil
}

// Start starts the finder.
func (finder *StaticFinder) Start() error {
	return nil
}

// Stop stops the finder.
func (finder *StaticFinder) Stop() error {
	return nil
}

// IsRunning returns true when the finder is running, otherwise false.
func (finder *StaticFinder) IsRunning() bool {
	return true
}

// String returns the description.
func (finder *StaticFinder) String() string {
	return FinderStatic
}
