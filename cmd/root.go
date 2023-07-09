package cmd

import (
	"os"

	"github.com/pvwnthem/gopwd/cmd/add"
	"github.com/pvwnthem/gopwd/cmd/vault"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopwd",
	Short: "A cli password manager written in go",
	Long:  "gopwd is an encrypted cli password manager (similar to password-store) written in golang",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(vault.VaultCmd)
	rootCmd.AddCommand(add.AddCmd)
}

func init() {
	addSubcommandPalettes()
}
