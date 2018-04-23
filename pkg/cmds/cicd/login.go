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

const loginPath string = "/user/login"

// NewCmdCicd creates a command for displaying the version of this binary
func NewCmdCicdLogin(name string) *cobra.Command {

	var (
		url             string
		username        string
		password        string
		defaultUsername string
		defaultURL      string
	)

	cmd := &cobra.Command{
		Use:   "login",
		Short: "CI/CD login command",
		Long:  "Run login command of Continuously Integration / Continuously Delivery",
		Run: func(cmd *cobra.Command, args []string) {
			//函数实现的功能
			//log.Debugf("[cicd] %s cicd login --url=%s --username=%s --password=%s\n", name, url, username, strings.Repeat("*", len(password)))
			conf := auth.ReadYaml()
			if len(conf.Hicli.Clusters) == 0 {
				//client.yml文件为空。让用户提供相关参数。获取并写入YAML
				url = auth.GetInput("Server")
				if ! auth.CheckUrl(url) {
					fmt.Println("Erro Server", url)
					return
				}
				username = auth.GetInput("Username")
				//password = auth.GetInput("Password")
			} else {
				//client.yml文件不为空，clusters下面有相关cluster信息。通过索引读取出来，并作为默认值待用
				lastIndex := conf.Hicli.LastIndex
				defaultURL = conf.Hicli.Clusters[lastIndex].Cluster
				defaultUsername = conf.Hicli.Clusters[lastIndex].Username
				url = auth.GetInput("Server[" + defaultURL + "]")
				if url == "" {
					url = defaultURL
					if ! auth.CheckUrl(url) {
						fmt.Println("Erro Server", url)
						return
					}
				}
				if username = auth.GetInput("Username[" + defaultUsername +  "]" ); username == "" {
					username = defaultUsername
				}
			}
			password = auth.GetInput("Password")
			userToken, err := auth.Login(url+loginPath, username, password)
			if err == nil && userToken != "" {
				err := auth.UpdateYAML(conf, url, username, userToken)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(username, "Login Sucess")
			} else {
				//提供的用户名与密码可能有误或者网络不可达
				fmt.Println("Login Failed ", err)
			}
		}, //函数功能实现结束标签
	}
	//pf := cmd.PersistentFlags()
	//pf.StringVarP(&url, "URL","u", "", "--url=http://www.example.com/")
	return cmd
}
