package cmds

import (
	"github.com/spf13/cobra"
	"fmt"
)

// NewCmdVersion creates a command for displaying the version of this binary
func NewCmdVersion(fullName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display client and server versions",
		Long:  "Display client and server versions.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s v2.0.0\n", fullName)
		},
	}

	return cmd
}
