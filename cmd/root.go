package cmd

import (
	"os"

	"github.com/amrnt/create-go-app/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	config.DefaultApp.Cmd = &cobra.Command{
		Use:   "create-go-app",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			return serverCmd.RunE(cmd, args)
		},
	}
}

// Execute ...
func Execute() {
	if c, err := config.DefaultApp.Cmd.ExecuteC(); err != nil {
		c.Println("")
		c.Println(c.UsageString())
		os.Exit(-1)
	}
}
