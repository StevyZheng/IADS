package disk

import "github.com/spf13/cobra"

var DiskCmd = &cobra.Command{
	Use:   "disk",
	Short: "disk operation",
}

func init() {
	DiskCmd.AddCommand(listCmd)
}
