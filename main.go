package main

import (
	"path/filepath"
	"os"
	"github.com/hidevopsio/hicli/pkg/cli"
	"github.com/hidevopsio/hi/boot/pkg/log"
)

func init() {
	log.SetLevel("debug")
}

func main() {

	basename := filepath.Base(os.Args[0])
	log.Debugf("[main] basename: %s", basename)
	command := cli.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
