package config

import "github.com/Sirupsen/logrus"

func init() {
	DefaultApp.Logger = logrus.New()
	DefaultApp.Logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	DefaultApp.Logger.Level = logrus.DebugLevel
}
