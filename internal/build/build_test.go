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
	"github.com/stretchr/testify/require"
)

func TestGetBuild(t *testing.T) {
	b := build.GetBuild()
	require.Equal(t, runtime.Version(), b.GoVersion)
}

func TestBuildLongVersion(t *testing.T) {
	b := build.GetBuild()
	str := b.LongVersion()
	require.Contains(t, str, "Version: ")
	require.Contains(t, str, "Revision: ")
	require.Contains(t, str, "Built: ")
	require.Contains(t, str, "Platform: ")
	require.Contains(t, str, "Go Version: ")
}
