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

		fmt.Printf("Successfully created vault at %s", filepath.Join(Path, Name))

		return nil
	},
}

func init() {
	VaultCmd.MarkPersistentFlagRequired("name")
	VaultCmd.AddCommand(initCmd)
}
