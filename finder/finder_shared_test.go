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
	"testing"
)

func TestSharedFinder(t *testing.T) {
	nodes := setupTestFinderNodes()

	finder := NewSharedFinder().(*SharedFinder)
	for _, node := range nodes {
		err := finder.addNode(node)
		if err != nil {
			t.Error(err)
		}
	}

	err := finder.Start()
	if err != nil {
		t.Error(err)
		return
	}

	err = finderTest(finder)
	if err != nil {
		t.Error(err)
	}

	err = finder.Stop()
	if err != nil {
		t.Error(err)
	}
}
