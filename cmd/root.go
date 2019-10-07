package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "iads",
	Short: "Roycom tools.",
	Long:  "Roycom tools.\nHa Ha Ha,I'm coming!",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
