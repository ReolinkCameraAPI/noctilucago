package cmd

import (
	"fmt"
	"github.com/ReolinkCameraAPI/noctilucago/config"
	"github.com/spf13/cobra"
	"os"
)

var (
	configPath = ""
	rootCmd = &cobra.Command{
		Use:              "app",
		TraverseChildren: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config.LoadConfigs(configPath)
		},
	}
)

func init() {
	rootCmd.VersionTemplate()
	rootCmd.HelpTemplate()
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "-c /path/to/file.yaml")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
