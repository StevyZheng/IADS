package cmd

import (
	"github.com/spf13/cobra"
	"iads/util"
)

func init() {
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	Run: func(cmd *cobra.Command, args []string) {
		util.ServerRunFunc()
	},
}
