package vault

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert [site] [flags]",
	Short: "Inserts a new password into the vault",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		site := args[0]
		vaultPath := filepath.Join(Path, Name)

		// Check if the vault exists
		_, err := os.Stat(vaultPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("vault does not exist")
		} else if err != nil {
			return fmt.Errorf("failed to check vault existence: %w", err)
		}

		// Create the directory
		dirPath := filepath.Join(vaultPath, site)
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		fmt.Printf("Inserted password for %s at %s", site, dirPath)

		return nil
	},
}

func init() {
	VaultCmd.AddCommand(insertCmd)
}
