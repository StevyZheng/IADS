package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = `version: 1.0.0
1. Compatible with all linux,include ubuntu.
2. Add version change description.
3. Add CentOS reboot test module.`

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of iads",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
