package vault

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [gpg-id]",
	Short: "Initializes a new vault at the specified path",
	Long:  "",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		vaultPath := Path

		// Check if the vault already exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultExistence, err)
		}
		if vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesExist, vaultPath)
		}

		err = util.CreateDirectory(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrDirectoryCreation, err)
		}

		err = util.CreateFile(filepath.Join(vaultPath, ".gpg-id"))
		if err != nil {
			return fmt.Errorf("failed to create gpg-id file: %w", err)
		}

		err = util.WriteToFile(filepath.Join(vaultPath, ".gpg-id"), args[0])
		if err != nil {
			return fmt.Errorf("failed to write to gpg-id file: %w", err)
		}

		err = util.CreateFile(filepath.Join(vaultPath, ".vault"))
		if err != nil {
			return fmt.Errorf("failed to create .vault file")
		}

		fmt.Printf("Successfully created vault at %s", vaultPath)

		return nil
	},
}

func init() {
	VaultCmd.MarkPersistentFlagRequired("name")
	VaultCmd.AddCommand(initCmd)
}
