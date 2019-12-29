package log

import (
	"github.com/spf13/cobra"
	"iads/util"
)

var printMessagesLogCmd = &cobra.Command{
	Use:   "messages",
	Short: "Print the messages error",
	Run: func(cmd *cobra.Command, args []string) {
		j, _ := util.MessagesErrorJson()
		println(j)
	},
}

var printMessagesLogTableCmd = &cobra.Command{
	Use:   "messagesTable",
	Short: "Print the messages error",
	Run: func(cmd *cobra.Command, args []string) {
		_ = util.MessagesErrorPrint()
	},
}
