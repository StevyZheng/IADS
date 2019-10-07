package cpu

import "github.com/spf13/cobra"

var CpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "cpu operation",
}

func init() {
	CpuCmd.AddCommand(getCpuInfoCmd)
	CpuCmd.AddCommand(burnCmd)
}
