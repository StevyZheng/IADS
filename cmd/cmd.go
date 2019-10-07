package cmd

import (
	"iads/cmd/cpu"
	"iads/cmd/disk"
	"iads/cmd/mb"
	"iads/cmd/zfs"
)

func init() {
	RootCmd.AddCommand(disk.DiskCmd)
	RootCmd.AddCommand(cpu.CpuCmd)
	RootCmd.AddCommand(mb.MbCmd)
	RootCmd.AddCommand(zfs.ZfsCmd)
}
