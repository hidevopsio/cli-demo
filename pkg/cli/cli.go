// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

package cli

import (
	"github.com/hidevopsio/hiboot/pkg/log"
	"runtime"
	"strings"
	"os"
	"fmt"
	"io"
	"github.com/spf13/cobra"
	"github.com/openshift/origin/pkg/cmd/util/term"
	"github.com/hidevopsio/hicli/pkg/cmds"
	"github.com/hidevopsio/hicli/pkg/cmds/cicd"
	"github.com/hidevopsio/hicli/pkg/common"
)

const productName = `hi`

type Commmands []*cobra.Command

var (
	cliLong = `
    ` + productName + ` client

    This client helps you develop, build, deploy, and run your applications on any
    OpenShift or Kubernetes compatible platform.`

	cliExplain = `
    To create a new application, login to your server and then change dir to your application:

        %[1]s login https://mycluster.mycompany.com
        %[1]s cicd run

    This will deploy you application to openshift / kubernetes cluster.

    Once your application is deployed, use the status to check if it runs properly:

        %[1]s status

    You should now see the URL the application can be reached at.

    To see the full list of commands supported, run '%[1]s --help'.`

)


func init()  {

}

func parseEnvOptions() *common.EnvOptions  {

	workDir, err := os.Getwd()
	var dirs []string
	if nil == err {
		dirs = strings.SplitN(workDir, string(os.PathSeparator), -1)
		log.Debug("[cli] dirs: ", dirs)
	}

	var app, project string

	l := len(dirs)
	if l > 1 {
		app = dirs[l - 1]
	}
	if l > 2 {
		project = dirs[l - 2]
	}

	envOptions := common.EnvOptions{
		App: app,
		Project: project,
	}

	return &envOptions
}

func NewCommandCLI(name, fullName string, in io.Reader, out, errout io.Writer) *cobra.Command {
	log.Debugf("name: %v", name)
	// Main command
	cmd := &cobra.Command{
		Use:   name,
		Short: "Command line tools hi-cli",
		Long:  "Command line tools for managing applications",
		Run: func(c *cobra.Command, args []string) {
			explainOut := term.NewResponsiveWriter(out) // TODO: term depends on openshift and kubernetes, lots of dependencies ...
			c.SetOutput(explainOut)
			fmt.Fprintf(explainOut, "%s\n\n%s\n", cliLong, fmt.Sprintf(cliExplain, name))
		},
	}

	cmd.AddCommand(
		cmds.NewCmdVersion(fullName),
		cicd.NewCmdCicd(name, parseEnvOptions()),
	)

	return cmd
}

func CommandFor(basename string) *cobra.Command {
	var cmd *cobra.Command

	in, out, errOut := os.Stdin, os.Stdout, os.Stderr

	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	cmd = NewCommandCLI(basename, "hi-cli", in, out, errOut)

	log.Debugf("[cli] basename: %s\n", basename)

	return cmd
}
