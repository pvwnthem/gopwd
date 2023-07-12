package vault

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes an existing vault at the specified path",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {
		vaultPath := filepath.Join(Path, Name)

		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultExistence, err)
		}
		if !vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		action, err := util.ConfirmAction()
		if err != nil {
			return fmt.Errorf(constants.ErrActionConfirm, err)
		}

		if action {
			err = util.RemoveDirectory(vaultPath)
			if err != nil {
				return fmt.Errorf("failed to remove vault: %w", err)
			}
		}

		fmt.Printf("Successfully removed vault at %s", vaultPath)

		return nil
	},
}

func init() {
	VaultCmd.MarkPersistentFlagRequired("name")
	VaultCmd.AddCommand(removeCmd)
}
