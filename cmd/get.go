package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"iads/lib/linux/hardware"
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getCpuInfoCmd)
	getCmd.AddCommand(getMbInfoCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get info",
}

var getCpuInfoCmd = &cobra.Command{
	Use:   "cpuinfo",
	Short: "Print the cpu info",
	Run: func(cmd *cobra.Command, args []string) {
		cpuInfo := new(hardware.CpuHwInfo)
		cpuInfo.GetCpuHwInfo()
		fmt.Println("model:", cpuInfo.Model)
		fmt.Println("sockets:", cpuInfo.Count)
		fmt.Println("cores:", cpuInfo.CoreCount)
		fmt.Println("stepping:", cpuInfo.Stepping)
	},
}

var getMbInfoCmd = &cobra.Command{
	Use:   "mbinfo",
	Short: "Print the motherborad info",
	Run: func(cmd *cobra.Command, args []string) {
		mbInfo := new(hardware.MotherboradInfo)
		mbInfo.GetMbInfo()
		fmt.Println("model:", mbInfo.Model)
		fmt.Println("biosVer:", mbInfo.BiosVer)
		fmt.Println("biosDate:", mbInfo.BiosDate)
	},
}
