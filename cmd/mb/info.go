package mb

import (
	"github.com/spf13/cobra"
	"iads/util"
)

var getMbInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the motherborad info",
	Run: func(cmd *cobra.Command, args []string) {
		util.MbInfoPrint()
	},
}
