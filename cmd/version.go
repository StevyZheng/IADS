package cmd

import (
	"github.com/spf13/cobra"
	"iads/util"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of iads",
	Run: func(cmd *cobra.Command, args []string) {
		util.VersionPrint()
	},
}
