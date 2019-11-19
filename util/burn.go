package util

import (
	"iads/lib/linux/hardware"
	. "iads/lib/logging"
	"iads/lib/math"
	"runtime"
	"sync"
)

//CPU Burn
func cpuBurnFunc(wg *sync.WaitGroup) {
	math.Gaos()
	(*wg).Done()
}

func CpuBurn() {
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
		go cpuBurnFunc(&wg)
	}
	wg.Wait()
}

//Disk Burn
