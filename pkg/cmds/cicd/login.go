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
	"github.com/hidevopsio/hi/boot/pkg/log"
	"strings"
)

// NewCmdCicd creates a command for displaying the version of this binary
func NewCmdCicdLogin(name string) *cobra.Command {

	var (
		url string
		username string
		password string
	)

	cmd := &cobra.Command{
		Use:   "login",
		Short: "CI/CD login command",
		Long:  "Run login command of Continuously Integration / Continuously Delivery",
		Run: func(cmd *cobra.Command, args []string) {
			log.Debugf("[cicd] %s cicd login --url=%s --username=%s --password=%s\n", name, url, username, strings.Repeat("*", len(password)))
		},
	}

	return cmd
}
