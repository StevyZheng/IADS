package mb

import "github.com/spf13/cobra"

var MbCmd = &cobra.Command{
	Use:   "mb",
	Short: "motherboard operation",
}

func init() {
	getOobCmd.Flags().String("bmcMac", "", "BMC mac address.")
	MbCmd.AddCommand(getMbInfoCmd)
	MbCmd.AddCommand(getOobCmd)
}
