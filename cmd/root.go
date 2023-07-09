package cmd

import (
	"os"
	"path/filepath"

	"github.com/pvwnthem/gopwd/cmd/vault"
	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var (
	Path string
	Name string
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
}

func init() {
	addSubcommandPalettes()
	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", filepath.Join(util.GetHomeDir(), ".gopwd"), "The path of the vault")
	rootCmd.PersistentFlags().StringVarP(&Name, "name", "n", "vault", "The name of the vault")

}
