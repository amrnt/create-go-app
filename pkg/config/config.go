package config

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// DefaultApp ...
	DefaultApp *App
)

// App ...
type App struct {
	Path    string
	Cmd     *cobra.Command
	Logger  *logrus.Logger
	Configs map[string]*viper.Viper
}

func init() {
	DefaultApp = &App{}
	// DefaultApp.Path =

	DefaultApp.Configs = make(map[string]*viper.Viper)
	DefaultApp.Configs["DefaultApp"] = viper.New()
}
