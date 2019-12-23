package util

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"iads/lib/common"
	"iads/lib/linux/hardware"
	"iads/lib/logging"
	"os"
)

func VersionPrint() {
	fmt.Println(version)
}

func CpuInfoPrint() {
	cpuInfo := new(hardware.CpuHwInfo)
	common.CheckError(cpuInfo.GetCpuHwInfo())
	fmt.Println("model:", cpuInfo.Model)
	fmt.Println("sockets:", cpuInfo.Count)
	fmt.Println("cores:", cpuInfo.CoreCount)
	fmt.Println("stepping:", cpuInfo.Stepping)
}

func MbInfoPrint() {
	mbInfo := new(hardware.MotherboradInfo)
	mbInfo.GetMbInfo()
	fmt.Println("model:", mbInfo.Model)
	fmt.Println("biosVer:", mbInfo.BiosVer)
	fmt.Println("biosDate:", mbInfo.BiosDate)
}

func DiskListPrint() {
	t := table.NewWriter()
	alignT := []text.Align{text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter, text.AlignCenter}
	t.SetOutputMirror(os.Stdout)
	t.SetAlignHeader(alignT)
	t.SetAlign(alignT)
	t.Style().Options.SeparateRows = true
	t.Style().Box = table.StyleBoxBold
	t.SetAutoIndex(true)

	t.AppendHeader(table.Row{"DEV", "MODEL", "SN", "WWN", "SIZE", "TYPE"})
	disks, err := hardware.Disk{}.DiskList()
	if err != nil {
		logging.FatalPrintln("Disk list error.")
		return
	}
	for _, disk := range disks {
		t.AppendRow(table.Row{disk.DevName, disk.Model, disk.Serial, disk.Wwn, fmt.Sprintf("%dGB", disk.Size), disk.DevType})
	}
	t.Render()
}

func DiskSmartInfoFromHBA() {

}
