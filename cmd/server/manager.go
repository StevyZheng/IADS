package server

import (
	"github.com/spf13/cobra"
)

var startManagerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Start manager server.",
	Run: func(cmd *cobra.Command, args []string) {
		println("manager api server is running...")
		//defer database.DBE.Close()

	},
}
