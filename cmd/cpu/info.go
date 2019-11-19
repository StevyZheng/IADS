package cpu

import (
	"github.com/spf13/cobra"
	"iads/util"
)

var getCpuInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the cpu info",
	Run: func(cmd *cobra.Command, args []string) {
		util.CpuInfoPrint()
	},
}
