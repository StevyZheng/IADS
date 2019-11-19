package cpu

import (
	"github.com/spf13/cobra"
	"iads/util"
)

var burnCmd = &cobra.Command{
	Use:   "burn",
	Short: "burn cpu: 100%",
	Run: func(cmd *cobra.Command, args []string) {
		util.CpuBurn()
	},
}
