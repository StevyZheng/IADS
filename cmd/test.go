package cmd

import (
	"github.com/spf13/cobra"
	"iads/lib/filesystem/zfs"
	"iads/lib/linux/hardware"
	"iads/lib/logging"
)

func init() {
	RootCmd.AddCommand(testCmd)
	testCmd.AddCommand(commonCmd)
	testCmd.AddCommand(runCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run roycom server test",
}

var commonCmd = &cobra.Command{
	Use:   "logging",
	Short: "test",
	Run: func(cmd *cobra.Command, args []string) {
		logging.FatalPrintln("ooooopppp")
		println("mimimimi")
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run roycom initer server test",
	Run: func(cmd *cobra.Command, args []string) {
		//n := hardware.NetInfo{}
		//_ = n.NetInit()
		//_ = common.DatasetCreate()
		disks, err := hardware.Disk{}.DiskList()
		if err != nil {
			return
		}

		pool := zfs.Pool{}
		pool.PoolName = "rpool"
		pool.Level = "raidz"
		pool.AddDiskList(disks)

		err = pool.MakeParts()
		if err != nil {
			println("part:", err.Error())
			return
		}

		//pool.Print()

		err = pool.Create(true)
		if err != nil {
			println("create:", err.Error())
			return
		}
	},
}
