package hardware

import (
	"github.com/jaypipes/ghw"
	"iads/lib/common"
	"iads/lib/stringx"
	"log"
	"path/filepath"
	"regexp"
)

type Disk struct {
	DevName  string
	Wwn      string
	Model    string
	Firmware string
	Serial   string
	Size     uint64
	DevType  string
}

func (d Disk) DiskList() (disks []Disk, err error) {
	block, err := ghw.Block()
	if err != nil {
		return
	}
	var diskT Disk
	for _, disk := range block.Disks {
		if IsCanUseDisk(disk.Name) {
			diskT = Disk{}
			diskT.DevName = disk.Name
			diskT.Model = disk.Model
			diskT.Serial = disk.SerialNumber
			diskT.Wwn = disk.WWN
			diskT.Size = disk.SizeBytes / 1024 / 1024 / 1024
			diskT.DevType = disk.Vendor
			disks = append(disks, diskT)
		}
	}
	return
}

func (d Disk) Partition() {
	//disk, _ := diskfs.Open(fmt.Sprintf("/dev/%s", d.DevName))
}

func ParseMountL() (osDisk []string) {
	ret, _ := common.ExecShellLinux("mount -l")
	ret1 := stringx.SearchSplitStringColumn(ret, ".+ / .+", " ", 1)
	for _, val := range ret1 {
		if stringx.MatchStr(val, ".+mapper.+") {
			pvRet, _ := common.ExecShellLinux("pvs --noheadings|awk -F'[0-9]' '{print$1}'")
			pvRet = filepath.Base(pvRet)
			osDisk = append(osDisk, pvRet)
		}
		if stringx.MatchStr(val, "/dev/sd.+") {
			reg := regexp.MustCompile("sd[a-z]+")
			ret := reg.FindStringSubmatch(val)
			if len(ret) != 0 {
				osDisk = append(osDisk, ret[0])
			}
		}
	}
	return osDisk
}

func GetOSDisk() (devName string) {
	if devNameT, err := common.ExecShellLinux("df|grep -P '/$'|awk '{print$1}'|awk -F'1' '{print$1}'|awk -F'/' '{print$3}'"); err != nil {
		log.Println("error: ", err.Error())
		return ""
	} else {
		return devNameT
	}
}

func IsCanUseDisk(devName string) (b bool) {
	osDevs := ParseMountL()
	for _, val := range osDevs {
		if val == "" {
			return false
		}
		if val == devName {
			return false
		}
		if stringx.MatchStr(devName, "sd.+") {
			return true
		} else {
			return false
		}
	}
	return false
}
