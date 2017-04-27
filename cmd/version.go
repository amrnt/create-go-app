package cmd

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/amrnt/create-go-app/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	config.DefaultApp.Cmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if config.BuildDate == "" {
			config.BuildDate = time.Now().Format(time.RFC3339)
		} else {
			t, _ := time.Parse("2006-01-02T15:04:05-0700", config.BuildDate)
			config.BuildDate = t.Format(time.RFC3339)
		}

		if config.CommitHash == "" {
			fmt.Printf("\033[32mcreate-go-app\033[0m v%s %s/%s BuildDate: %s\n", config.Version, runtime.GOOS, runtime.GOARCH, config.BuildDate)
		} else {
			fmt.Printf("\033[32mcreate-go-app\033[0m v%s-%s %s/%s BuildDate: %s\n", config.Version, strings.ToUpper(config.CommitHash), runtime.GOOS, runtime.GOARCH, config.BuildDate)
		}
	},
}
