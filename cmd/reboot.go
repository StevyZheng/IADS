package cmd

import (
	"github.com/spf13/cobra"
	"iads/util"
)

func init() {
	RootCmd.AddCommand(RebootCmd)
}

var RebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "reboot test in linux, for ZStack now.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := util.RebootFunc(); err != nil {
			println(err.Error())
			return
		}
	},
}
