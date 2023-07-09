package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Path string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new vault at the specified path",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	initCmd.Flags().StringVarP(&Path, "path", "p", "", "The path to create the vault at")

	if err := initCmd.MarkFlagRequired("path"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(initCmd)
}
