package cmd

import (
	"github.com/spf13/cobra"
	"iads/server/internals/app/routers"
	"iads/server/internals/pkg/models/sys"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	Run: func(cmd *cobra.Command, args []string) {
		println("iads api server is running...")
		//defer database.DBE.Close()
		sys.ModelInit()
		router := routers.InitRouter()
		_ = router.Run(":80")
	},
}
