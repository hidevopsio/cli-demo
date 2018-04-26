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

package cicd

import (
	"github.com/spf13/cobra"
	"github.com/hidevopsio/hicli/pkg/common"
	"github.com/hidevopsio/hiboot/pkg/log"
	"github.com/hidevopsio/hicli/pkg/services/cicd"
	"fmt"
)

var (
	profile string
	project string
	app     string
	name    string
)

// NewCmdCicd creates a command for displaying the version of this binary
func NewCmdCicdRun(name string, envOptions *common.EnvOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "CI/CD run command",
		Long:  "Run run command of Continuously Integration / Continuously Delivery",
		Run: func(cmd *cobra.Command, args []string) {
			log.Debugf("[cicd] %s cicd run --profile=%s --project=%s --app=%s\n", name, profile, project, app)
			token, url := cicd.GetTokenUrl()
			if token == "" || url == "" {
				fmt.Println("No Login,Plase run hicli cicd login first!")
				return
			}
			env, err := cicd.InitEnvOpt(name, profile, app, project)
			if err == nil {
				if err = cicd.CICDRun(url, token, env); err == nil {
					fmt.Printf("App %s Deploy Sucess\n", app)
				} else {
					fmt.Printf("App %s Deploy Failed.%s\n", app, err.Error())
				}
			} else {
				fmt.Printf("App %s Deploy Failed.%s\n", app, err.Error())
			}
		},
	}

	pf := cmd.PersistentFlags()

	pf.StringVarP(&profile, "profile", "p", "dev", "--profile=test")
	pf.StringVarP(&project, "project", "P", "", "--profile=project-name")
	pf.StringVarP(&app, "app", "a", "", "--app=my-app")
	pf.StringVarP(&name, "name", "n", "", "--name=my-name")

	return cmd
}
