package cpu

import (
	"github.com/spf13/cobra"
	"iads/lib/linux/hardware"
	. "iads/lib/logging"
	"iads/lib/math"
	"runtime"
	"sync"
)

func burnFunc(wg *sync.WaitGroup) {
	math.Gaos()
	(*wg).Done()
}

var burnCmd = &cobra.Command{
	Use:   "burn",
	Short: "burn cpu: 100%",
	Run: func(cmd *cobra.Command, args []string) {
		cpuInfo := hardware.CpuHwInfo{}
		_ = cpuInfo.GetCpuHwInfo()
		if cpuInfo.CoreCount <= 0 {
			FatalPrintln("getCpuInfo error.")
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
