package zfs

import "github.com/spf13/cobra"

var ZfsCmd = &cobra.Command{
	Use:   "libzfs-tmp",
	Short: "libzfs-tmp operation",
}

func init() {
	ZfsCmd.AddCommand(getZfsVersionCmd)
}
