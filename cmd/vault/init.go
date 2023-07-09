package vault

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new vault at the specified path",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {
		vaultPath := filepath.Join(Path, Name)

		// Check if the vault already exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to check vault existence: %w", err)
		}
		if vaultExists {
			return errors.New("vault already exists")
		}

		err = util.CreateDirectory(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		fmt.Printf("Successfully created vault at %s", vaultPath)

		return nil
	},
}

func init() {
	VaultCmd.MarkPersistentFlagRequired("name")
	VaultCmd.AddCommand(initCmd)
}
