package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "cp [password] [new-password]",
	Short: "Copies a password from one site to another",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		site := args[0]
		destination := args[1]

		vaultPath := Path
		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultExistence, err)
		}
		if !vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		// Check if the password exists
		passwordPath := filepath.Join(vaultPath, site, "password")
		passwordExists, err := util.Exists(passwordPath)
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordExistence, err)
		}
		if !passwordExists {
			return fmt.Errorf(constants.ErrPasswordDoesNotExist)
		}

		var action bool

		if !Force {
			action, err = util.ConfirmAction()
			if err != nil {
				return fmt.Errorf(constants.ErrActionConfirm, err)
			}
		} else {
			action = true
		}

		if action {
			// copy the password to the new site
			err := util.CreateDirectory(filepath.Join(vaultPath, destination))
			if err != nil {
				return fmt.Errorf("failed to create new directory: %w", err)
			}

			err = util.CopyFile(filepath.Join(vaultPath, site, "password"), filepath.Join(vaultPath, destination, "password"))
			if err != nil {
				return fmt.Errorf("failed to copy password: %w", err)
			}

			fmt.Printf("Copied password from %s to %s \n", site, destination)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
}
