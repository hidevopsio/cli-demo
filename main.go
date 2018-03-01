package main

import (
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"os"
	"github.com/hi-devops-io/hi-cli/pkg/cli"
)

func init() {

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.SetLevel(log.DebugLevel)

	log.Debugln("[main] init()")
}

func main() {

	basename := filepath.Base(os.Args[0])
	log.Debugf("[main] basename: %s", basename)
	command := cli.CommandFor(basename)
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
