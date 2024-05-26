// Copyright 2024 Gustavo Salomão
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build_test

import (
	"runtime"
	"testing"

	"github.com/gsalomao/terracov/internal/build"
	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	info := build.GetInfo()
	assert.Equal(t, runtime.Version(), info.GoVersion)
}

func TestInfoLongVersion(t *testing.T) {
	info := build.GetInfo()
	assert.NotEmpty(t, info.LongVersion())
}
