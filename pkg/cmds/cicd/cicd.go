package cicd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/hidevopsio/hicli/pkg/common"
)

// NewCmdCicd creates a command for displaying the version of this binary
func NewCmdCicd(name string, envOptions *common.EnvOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cicd",
		Short: "CI/CD command",
		Long:  "Run cicd command Continuously Integration / Continuously Delivery ",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("\n%s cicd\n", name)
		},
	}

	cmd.AddCommand(NewCmdCicdRun(name, envOptions))

	return cmd
}
