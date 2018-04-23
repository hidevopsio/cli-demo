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

package main

import (
	"path/filepath"
	"os"
	"github.com/hidevopsio/hicli/pkg/cli"
	"github.com/hidevopsio/hiboot/pkg/log"
	"fmt"
)


func init() {
	log.SetLevel("debug")
}

func main() {
	basename := filepath.Base(os.Args[0])
	//log.Debugf("[main] basename: %s", basename)
	fmt.Println("basename: ",basename)
	command := cli.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
