package vault

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an existing vault at the specified path",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {
		vaultPath := filepath.Join(Path, Name)

		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to check vault existence: %w", err)
		}
		if !vaultExists {
			return errors.New("vault does not exist")
		}

		err = util.RemoveDirectory(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to remove vault: %w", err)
		}

		fmt.Printf("Successfully removed vault at %s", vaultPath)

		return nil
	},
}

func init() {
	VaultCmd.MarkPersistentFlagRequired("name")
	VaultCmd.AddCommand(removeCmd)
}
