package cmd

import (
	"os"

	"github.com/amrnt/create-go-app/pkg/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "create-go-app",
	Short: "",
	Long:  ``,
}

func init() {
	config.DefaultApp.Cmd = rootCmd
}

// Execute ...
func Execute() {
	if err := config.DefaultApp.Cmd.Execute(); err != nil {
		config.DefaultApp.Logger.Error(err)
		os.Exit(-1)
	}
}
