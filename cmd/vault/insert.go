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
		//check if the vault exists

		if _, err := os.Stat(filepath.Join(Path, Name)); os.IsNotExist(err) {
			return fmt.Errorf("vault does not exist")
		}
		err := os.MkdirAll(filepath.Join(Path, Name, site), 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		return nil
	},
}

func init() {
	VaultCmd.AddCommand(insertCmd)
}
