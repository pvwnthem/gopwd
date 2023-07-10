package config

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Path string
	Name string
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Config is a palette that contains commands to manage the configuration",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := ConfigCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Set default values for flags
	ConfigCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "Path value")
	ConfigCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Name value")
}
