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
	"fmt"
	"github.com/hidevopsio/hicli/pkg/services/auth"
)

// NewCmdCicd creates a command for displaying the version of this binary
func NewCmdCicdLogin(name string) *cobra.Command {

	var (
		url string
		//username string
		//password string
	)

	cmd := &cobra.Command{
		Use:   "login",
		Short: "CI/CD login command",
		Long:  "Run login command of Continuously Integration / Continuously Delivery",
		Run: func(cmd *cobra.Command, args []string) {
			//函数实现的功能
			//log.Debugf("[cicd] %s cicd login --url=%s --username=%s --password=%s\n", name, url, username, strings.Repeat("*", len(password)))
			fmt.Println("The Func args is ", args)
			lastIndex,exists := auth.GetLastIndex()
			if !exists {
				url := auth.GetInput("url")
				fmt.Println(url)
			}
			fmt.Println(lastIndex,url)
		},//函数功能实现结束标签
	}
	pf := cmd.PersistentFlags()
	pf.StringVarP(&url, "URL","u", "", "--url=http://www.example.com/")
	return cmd
}
