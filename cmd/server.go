package cmd

import (
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
	RunE: func(cmd *cobra.Command, args []string) error {
		config.DefaultApp.Logger.Info("Running server...")
		return nil
	},
}
