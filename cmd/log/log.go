package log

import "github.com/spf13/cobra"

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "log operation",
}

func init() {
	LogCmd.AddCommand(printMessagesLogCmd)
	LogCmd.AddCommand(printMessagesLogTableCmd)
}
