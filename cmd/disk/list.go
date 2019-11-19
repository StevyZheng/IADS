package disk

import (
	"github.com/spf13/cobra"
	"iads/util"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the disk list",
	Run: func(cmd *cobra.Command, args []string) {
		util.DiskListPrint()
	},
}
