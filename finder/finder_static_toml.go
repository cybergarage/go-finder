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
	"github.com/BurntSushi/toml"
	"github.com/cybergarage/go-finder/finder/node"
	"github.com/cybergarage/go-logger/log"
)

// StaticFinder represents a simple static finder.
type StaticTOMLFinder struct {
	*StaticFinder
}

// NewStaticFinderWithConfig returns a new static finder with specified nodes.
func NewStaticFinderWithConfig(config FinderConfig) Finder {
	nodes := []Node{}
	for _, host := range config.Hosts {
		node := node.NewBaseNode()
		node.SetHost(host)
		nodes = append(nodes, node)
	}
	return NewStaticFinderWithNodes(nodes)
}

// NewStaticFinderWithTOML returns a new static finder with specified nodes.
func NewStaticFinderWithTOML(filename string) (Finder, error) {
	conf := Config{}
	if filename != "" {
		log.Tracef("TOML Config file path: %s", filename)
		_, err := toml.DecodeFile(filename, &conf)
		if err != nil {
			return nil, err
		}
		log.Tracef("Got config: %s", filename)
	}
	return NewStaticFinderWithConfig(conf.Finder), nil
}

// String returns the description
func (finder *StaticTOMLFinder) String() string {
	return FinderStaticToml
}
