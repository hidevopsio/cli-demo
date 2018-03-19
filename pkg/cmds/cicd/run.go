package cicd


import (
	"github.com/spf13/cobra"
	"github.com/hidevopsio/hicli/pkg/common"
	"github.com/hidevopsio/hi/boot/pkg/log"
)

var (
	profile string
	project string
	app string
)

// NewCmdCicd creates a command for displaying the version of this binary
func NewCmdCicdRun(name string, envOptions *common.EnvOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "CI/CD run command",
		Long:  "Run run command of Continuously Integration / Continuously Delivery",
		Run: func(cmd *cobra.Command, args []string) {
			if "" == app {
				app = envOptions.App
			}
			if "" == project {
				project = envOptions.Project
			}

			log.Debugf("[cicd] %s cicd run --profile=%s --project=%s --app=%s\n", name, profile, project, app)
		},
	}

	cmd.PersistentFlags().StringVarP(&profile, "profile","p", "dev", "--profile=test")
	cmd.PersistentFlags().StringVarP(&project, "project","P", "", "--profile=project-name")
	cmd.PersistentFlags().StringVarP(&app, "app","a", "", "--app=my-app")

	return cmd
}
