package cmd

import (
	"fmt"

	"github.com/amrnt/create-go-app/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	config.DefaultApp.Cmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running server...")
	},
}
