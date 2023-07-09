package vault

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new vault at the specified path",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the vault already exists
		if _, err := os.Stat(filepath.Join(Path, Name)); !os.IsNotExist(err) {
			return errors.New("vault already exists")
		}

		err := os.MkdirAll(filepath.Join(Path, Name), 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		return nil
	},
}

func init() {
	homeDir, homeDirErr := os.UserHomeDir()

	if homeDirErr != nil {
		panic(homeDirErr)
	}

	initCmd.Flags().StringVarP(&Path, "path", "p", filepath.Join(homeDir, ".gopwd"), "The path to create the vault at")
	initCmd.Flags().StringVarP(&Name, "name", "n", "", "The name of the vault")

	VaultCmd.AddCommand(initCmd)

	err := initCmd.MarkFlagRequired("name")
	if err != nil {
		panic(err)
	}
}
