package vault

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an existing vault at the specified path",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the vault exists
		if _, err := os.Stat(filepath.Join(Path, Name)); os.IsNotExist(err) {
			return errors.New("vault does not exist")
		}

		err := os.RemoveAll(filepath.Join(Path, Name))
		if err != nil {
			return fmt.Errorf("failed to remove vault: %w", err)
		}

		fmt.Printf("Successfully removed vault at %s", filepath.Join(Path, Name))

		return nil
	},
}

func init() {
	VaultCmd.AddCommand(removeCmd)
}
