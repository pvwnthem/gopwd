package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var (
	Recursive bool
)

var rmCmd = &cobra.Command{
	Use:   "rm [site]",
	Short: "Removes a password from the vault",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		site := args[0]
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

		if !Recursive {
			passwordPath := filepath.Join(vaultPath, site, "password")
			passwordExists, err := util.Exists(passwordPath)
			if err != nil {
				return fmt.Errorf(constants.ErrPasswordExistence, err)
			}
			if !passwordExists {
				return fmt.Errorf(constants.ErrPasswordDoesNotExist)
			}
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

		// Remove the password directory

		passwordParentFolderPath := filepath.Join(vaultPath, site)

		if action {
			nestedDirs, err := util.GetNestedDirectories(passwordParentFolderPath)
			if err != nil {
				return fmt.Errorf("failed to get nested directories: %w", err)
			}

			if Recursive || len(nestedDirs) == 0 {
				err = util.RemoveDirectory(passwordParentFolderPath)
				if err != nil {
					return fmt.Errorf("failed to remove password directory: %w", err)
				}
			} else {
				if len(nestedDirs) > 0 {
					err = util.RemoveFile(filepath.Join(passwordParentFolderPath, "password"))
					if err != nil {
						return fmt.Errorf("failed to remove password file: %w", err)
					}
				} else {
					err = util.RemoveDirectory(passwordParentFolderPath)
					if err != nil {
						return fmt.Errorf("failed to remove password directory: %w", err)
					}
				}
			}

			fmt.Println("done")
		}

		return nil
	},
}

func init() {

	rmCmd.Flags().BoolVarP(&Recursive, "recursive", "r", false, "remove the password directory and all items inside of it")

	rootCmd.AddCommand(rmCmd)
}
