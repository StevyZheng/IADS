package mb

import (
	"fmt"
	"github.com/spf13/cobra"
	"iads/lib/linux/hardware"
)

var getMbInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the motherborad info",
	Run: func(cmd *cobra.Command, args []string) {
		mbInfo := new(hardware.MotherboradInfo)
		mbInfo.GetMbInfo()
		fmt.Println("model:", mbInfo.Model)
		fmt.Println("biosVer:", mbInfo.BiosVer)
		fmt.Println("biosDate:", mbInfo.BiosDate)
	},
}
