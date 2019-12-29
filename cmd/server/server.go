package server

import "github.com/spf13/cobra"

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
}

func init() {
	ServerCmd.AddCommand(startIadsCmd)
	ServerCmd.AddCommand(startManagerCmd)
}
