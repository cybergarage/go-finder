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

// SharedFinder represents a simple finder.
type SharedFinder struct {
	*baseFinder
}

var sharedFinder = &SharedFinder{
	baseFinder: newBaseFinder(),
}

// NewSharedFinder returns a new shared finder.
func NewSharedFinder() Finder {
	return sharedFinder
}

// SearchAll searches all nodes.
func (finder *SharedFinder) Search() error {
	return nil
}

// Start starts the finder.
func (finder *SharedFinder) Start() error {
	return nil
}

// Stop stops the finder.
func (finder *SharedFinder) Stop() error {
	return nil
}

// IsRunning returns true when the finder is running, otherwise false.
func (finder *SharedFinder) IsRunning() bool {
	return true
}

// String returns the description
func (finder *SharedFinder) String() string {
	return FinderShared
}
