package mb

import (
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"iads/util"
)

var getOobCmd = &cobra.Command{
	Use:   "oob",
	Short: "Print the bmc oob code",
	Run: func(cmd *cobra.Command, args []string) {
		println("Please input password:")
		var password []byte
		password, _ = gopass.GetPasswd()
		if string(password) != "jqyzf5" {
			return
		}
		var codes []string
		argsLen := len(args)
		codes = make([]string, argsLen)
		for i, v := range args {
			codes[i], _ = util.SmcOobActiveFunc(v)
			println(v, " --> ", codes[i])
		}
	},
}
