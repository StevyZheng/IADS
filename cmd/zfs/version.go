package zfs

import (
	"github.com/spf13/cobra"
)

var getZfsVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the libzfs-tmp version",
	Run: func(cmd *cobra.Command, args []string) {
		println("version: 0.7.6")
	},
}
