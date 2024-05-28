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

package command

import (
	"io"

	"github.com/gsalomao/terracov/internal/build"
	"github.com/urfave/cli/v2"
)

// Version is the CLI command to show the version and build information.
var Version = &cli.Command{
	Name:            "version",
	Usage:           "Show version and build information",
	HideHelpCommand: true,
	Action: func(cCtx *cli.Context) error {
		return showVersion(cCtx.App.Writer)
	},
}

func showVersion(w io.Writer) error {
	b := build.GetBuild()
	_, err := w.Write([]byte(b.LongVersion()))
	return err
}
