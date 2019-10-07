package zfs

import "github.com/spf13/cobra"

var ZfsCmd = &cobra.Command{
	Use:   "zfs",
	Short: "zfs operation",
}

func init() {
	ZfsCmd.AddCommand(getZfsVersionCmd)
}
