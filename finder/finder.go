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
	"regexp"
)

// FinderSearchListener a listener for Finder.
type FinderSearchListener interface {
	FinderSearchResponseReceived(*Node)
}

// FinderNotifyListener a listener for Finder.
type FinderNotifyListener interface {
	FinderNotifyReceived(*Node)
}

// Finder represents an abstract interface.
type Finder interface {
	// SearchAll searches all nodes.
	Search() error
	// SetSearchListener sets a specified listener.
	SetSearchListener(FinderSearchListener) error
	// SetNotifyListener sets a specified listener.
	SetNotifyListener(FinderNotifyListener) error
	// GetAllNodes returns all found nodes.
	GetAllNodes() ([]Node, error)
	// GetPrefixNodes returns only nodes matching with a specified start string.
	GetPrefixNodes(string) ([]Node, error)
	// GetRegexpNodes returns only nodes matching with a specified regular expression.
	GetRegexpNodes(*regexp.Regexp) ([]Node, error)
	// GetNeighborhoodNode returns a neighborhood node of the specified node.
	GetNeighborhoodNode(node Node) (Node, error)
	// Start starts the finder.
	Start() error
	// Stop stops the finder.
	Stop() error
	// IsRunning returns true when the finder is running, otherwise false.
	IsRunning() bool
	// String returns the description
	String() string
}
