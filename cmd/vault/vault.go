package vault

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	Path string
	Name string
)

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Vault is a palette that contains commands to manage and create vaults",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	err := vaultCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	homeDir, homeDirErr := os.UserHomeDir()

	if homeDirErr != nil {
		panic(homeDirErr)
	}

	err := initCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}

	vaultCmd.PersistentFlags().StringVarP(&Path, "path", "p", filepath.Join(homeDir, ".gopwd"), "The path to create the vault at")
	vaultCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "The name of the vault")
}
