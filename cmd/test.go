package cmd

import (
	//"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"iads/lib/linux/hardware"
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(commonCmd)
	testCmd.AddCommand(runCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run roycom server test",
}

var commonCmd = &cobra.Command{
	Use:   "initer",
	Short: "test",
	Run: func(cmd *cobra.Command, args []string) {
		d := hardware.DmiInfo{}
		d.Run()
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run roycom initer server test",
	Run: func(cmd *cobra.Command, args []string) {
		n := hardware.NetInfo{}
		_ = n.NetInit()
	},
}
