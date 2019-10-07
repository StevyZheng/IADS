package mb

import "github.com/spf13/cobra"

var MbCmd = &cobra.Command{
	Use:   "mb",
	Short: "motherboard operation",
}

func init() {
	MbCmd.AddCommand(getMbInfoCmd)
}
