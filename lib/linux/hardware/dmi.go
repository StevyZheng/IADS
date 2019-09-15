package hardware

import (
	"fmt"
	"github.com/dselans/dmidecode"
	"github.com/shirou/gopsutil/cpu"
)

type DmiInfo struct {
}

func (e *DmiInfo) Run() {
	dmi := dmidecode.New()
	if err := dmi.Run(); err != nil {
		fmt.Printf("Unable to get dmidecode information. Error: %v\n", err)
	}
	for handle, record := range dmi.Data {
		fmt.Println("Checking record:", handle)
		//for k, v := range record {
		//	fmt.Printf("Key: %v Val: %v\n", k, v)
		//}
		fmt.Println(record[0])
	}
}

func (e *DmiInfo) Run2() {
	x, _ := cpu.Info()
	fmt.Println(len(x))
	for _, i := range x {
		j := cpu.InfoStat(i).ModelName
		fmt.Println(j)
	}
}
