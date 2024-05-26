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

package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func init() {
	cli.AppHelpTemplate = `{{.Usage}}{{if .Description}}

Description:
  {{.Description}}{{end}}

Usage:
  {{.HelpName}}{{if .VisibleFlags}} [options]{{end}}{{if .Commands}} [command]{{end}}{{if .VisibleCommands}}

Available Commands:{{template "visibleCommandTemplate" .}}{{end}}{{if .VisibleFlags}}

Options:{{template "visibleFlagTemplate" .}}{{end}}

Use "{{.Name}}{{if .Commands}} [command]{{end}} --help" for more information about a command.
`

	cli.CommandHelpTemplate = `{{.Usage}}{{if .Description}}

Description:
  {{.Description}}{{end}}

Usage:
   {{.HelpName}}{{if .VisibleFlags}} [options]{{end}}{{if .VisibleCommands}} [command]{{end}}{{if .VisibleCommands}}

Available Commands:{{template "visibleCommandTemplate" .}}{{end}}{{if .VisibleFlags}}

Options:{{template "visibleFlagTemplate" .}}{{end}}
`

	cli.SubcommandHelpTemplate = cli.CommandHelpTemplate
}

func main() {
	app := &cli.App{
		Name:  "terracov",
		Usage: "Static analysis scanner for Terraform",
		Description: "" +
			"Terracov scans cloud infrastructure provisioned using Terraform " +
			"and detects security issues and misconfiguration using standard " +
			"or custom rules.",
		HideHelpCommand: true,
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
