package config

import (
	"github.com/spf13/cobra"
)

var (
	// DefaultApp ...
	DefaultApp *App
)

// App ...
type App struct {
	Cmd *cobra.Command
}

func init() {
	DefaultApp = &App{}
}
