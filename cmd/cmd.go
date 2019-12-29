package cmd

import (
	"iads/cmd/cpu"
	"iads/cmd/disk"
	"iads/cmd/log"
	"iads/cmd/mb"
	"iads/cmd/server"
	"iads/cmd/zfs"
)

func init() {
	RootCmd.AddCommand(disk.DiskCmd)
	RootCmd.AddCommand(cpu.CpuCmd)
	RootCmd.AddCommand(mb.MbCmd)
	RootCmd.AddCommand(zfs.ZfsCmd)
	RootCmd.AddCommand(server.ServerCmd)
	RootCmd.AddCommand(log.LogCmd)
}
