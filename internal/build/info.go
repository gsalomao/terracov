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

package build

import (
	"bytes"
	"fmt"
	"runtime"
	"text/tabwriter"
)

// Variables updated at build time.
var (
	version   = "0.0.0"               // The application version.
	revision  = "0"                   // The commit ID of the build.
	buildTime = "2024-01-01 00:00:00" // The build time in UTC (year-month-day hour:min:sec).
)

// Info contains build information.
type Info struct {
	// The application version.
	Version string

	// The commit ID of the build.
	Revision string

	// The build time in UTC (year-month-day hour:min:sec).
	BuildTime string

	// The runtime platform (architecture and operating system).
	Platform string

	// The runtime Go version.
	GoVersion string
}

// GetInfo returns build information.
func GetInfo() Info {
	return Info{
		Version:   version,
		Revision:  revision,
		BuildTime: buildTime,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		GoVersion: runtime.Version(),
	}
}

// LongVersion returns the application version in long format.
func (b *Info) LongVersion() string {
	var buf bytes.Buffer
	tw := tabwriter.NewWriter(&buf, 2, 1, 2, ' ', 0)

	_, _ = fmt.Fprintf(tw, "Version:       %s\n", b.Version)
	_, _ = fmt.Fprintf(tw, "Revision:      %s\n", b.Revision)
	_, _ = fmt.Fprintf(tw, "Built:         %s\n", b.BuildTime)
	_, _ = fmt.Fprintf(tw, "Platform:      %s\n", b.Platform)
	_, _ = fmt.Fprintf(tw, "Go version:    %s\n", b.GoVersion)

	_ = tw.Flush()
	return buf.String()
}
