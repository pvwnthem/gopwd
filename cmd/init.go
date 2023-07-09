package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	Path string
	Name string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new vault at the specified path",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		// Check if the vault already exists
		if _, err := os.Stat(filepath.Join(Path, Name)); !os.IsNotExist(err) {
			panic("Vault already exists")
		}

		err := os.MkdirAll(filepath.Join(Path, Name), 0755)
		if err != nil {
			panic(err)
		}
	},
}

func init() {

	homeDir, homeDirErr := os.UserHomeDir()

	if homeDirErr != nil {
		panic(homeDirErr)
	}

	initCmd.Flags().StringVarP(&Path, "path", "p", filepath.Join(homeDir, ".gopwd"), "The path to create the vault at")
	initCmd.Flags().StringVarP(&Name, "name", "n", "", "The name of the vault")

	err := initCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(initCmd)
}
