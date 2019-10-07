package cpu

import (
	"fmt"
	"github.com/spf13/cobra"
	"iads/lib/common"
	"iads/lib/linux/hardware"
)

var getCpuInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the cpu info",
	Run: func(cmd *cobra.Command, args []string) {
		cpuInfo := new(hardware.CpuHwInfo)
		common.CheckError(cpuInfo.GetCpuHwInfo())
		fmt.Println("model:", cpuInfo.Model)
		fmt.Println("sockets:", cpuInfo.Count)
		fmt.Println("cores:", cpuInfo.CoreCount)
		fmt.Println("stepping:", cpuInfo.Stepping)
	},
}
