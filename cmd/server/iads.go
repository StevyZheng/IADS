package server

import (
	"github.com/spf13/cobra"
	"iads/server/iads/internals/app/routers"
	"iads/server/iads/internals/pkg/models/sys"
)

var startIadsCmd = &cobra.Command{
	Use:   "iads",
	Short: "Start iads server.",
	Run: func(cmd *cobra.Command, args []string) {
		println("iads api server is running...")
		//defer database.DBE.Close()
		sys.ModelInit()
		router := routers.InitRouter()
		_ = router.Run(":80")
	},
}
