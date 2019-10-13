package zfs

import (
	"fmt"
	"iads/lib/common"
	"iads/lib/linux/hardware"
)

type Pool struct {
	Disks    []hardware.Disk
	Level    string
	PoolName string
}

func (p Pool) Print() {
	for _, v := range p.Disks {
		println(v.PartID)
	}
}

func (p *Pool) AddDisk(disk hardware.Disk) {
	p.Disks = append(p.Disks, disk)
}

func (p *Pool) AddDiskList(diskList []hardware.Disk) {
	for i, _ := range diskList {
		p.Disks = append(p.Disks, diskList[i])
	}
}

func (p *Pool) MakeParts() (err error) {
	for i, _ := range p.Disks {
		err = p.Disks[i].MakePartition()
		if err != nil {
			return
		}
	}
	_ = hardware.Disk{}.PartProbe()
	for i, _ := range p.Disks {
		err = p.Disks[i].FillPartID()
		if err != nil {
			return
		}
	}
	return
}

func (p *Pool) Create(isForce bool) (err error) {
	Command := fmt.Sprintf("zpool create %s %s", p.PoolName, p.Level)
	for i, _ := range p.Disks {
		//p.Disks[i].Print()
		Command = fmt.Sprintf("%s %s", Command, p.Disks[i].PartID)
	}
	if isForce {
		Command = fmt.Sprintf("%s -f", Command)
	}
	//println(Command)
	_, err = common.ExecShellLinux(Command)
	return err
}

func (p *Pool) Destroy() (err error) {
	Command := fmt.Sprintf("zpool destroy %s", p.PoolName)
	_, err = common.ExecShellLinux(Command)
	return
}
