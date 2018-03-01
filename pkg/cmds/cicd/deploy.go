package cicd


import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/hi-devops-io/hi-cli/pkg/common"
)

var (
	profile string
	project string
	app string
)

// NewCmdCicd creates a command for displaying the version of this binary
func NewCmdCicdDeply(name string, envOptions *common.EnvOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "CI/CD deploy command",
		Long:  "Run deploy command of Continuously Integration / Continuously Delivery",
		Run: func(cmd *cobra.Command, args []string) {
			if "" == app {
				app = envOptions.App
			}
			if "" == project {
				project = envOptions.Project
			}

			log.Debugf("[cicd] %s cicd deploy --profile=%s --project=%s --app=%s\n", name, profile, project, app)
		},
	}

	cmd.PersistentFlags().StringVarP(&profile, "profile","p", "dev", "--profile=test")
	cmd.PersistentFlags().StringVarP(&project, "project","P", "", "--profile=project-name")
	cmd.PersistentFlags().StringVarP(&app, "app","a", "", "--app=my-app")

	return cmd
}
