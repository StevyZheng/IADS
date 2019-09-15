package cmd

import (
	"github.com/spf13/cobra"
	"iads/lib/linux/hardware"
	"iads/lib/math"
	"log"
	"runtime"
	"sync"
)

func init() {
	rootCmd.AddCommand(burnCmd)
}

func burnFunc(wg *sync.WaitGroup) {
	math.Gaos()
	(*wg).Done()
}

var burnCmd = &cobra.Command{
	Use:   "burn",
	Short: "burn cpu: 100%",
	Run: func(cmd *cobra.Command, args []string) {
		cpuInfo := hardware.CpuHwInfo{}
		cpuInfo.GetCpuHwInfo()
		if cpuInfo.CoreCount <= 0 {
			log.Fatal("getCpuInfo error.")
			return
		}
		runtime.GOMAXPROCS(cpuInfo.CoreCount)
		var wg sync.WaitGroup
		for i := 0; i < cpuInfo.CoreCount; i++ {
			wg.Add(1)
			go burnFunc(&wg)
		}
		wg.Wait()
	},
}
