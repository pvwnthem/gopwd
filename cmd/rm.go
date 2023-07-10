package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm [site]",
	Short: "Removes a password from the vault",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		site := args[0]
		vaultPath := filepath.Join(Path, Name)

		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to check vault existence: %w", err)
		}
		if !vaultExists {
			return fmt.Errorf("vault does not exist at %s", vaultPath)
		}

		// Check if the password exists
		passwordPath := filepath.Join(vaultPath, site)
		passwordExists, err := util.Exists(passwordPath)
		if err != nil {
			return fmt.Errorf("failed to check password existence: %w", err)
		}
		if !passwordExists {
			return fmt.Errorf("password does not exist")
		}

		action, err := util.ConfirmAction()
		if err != nil {
			return fmt.Errorf("failed to confirm action: %w", err)
		}

		// Remove the password directory
		if action {
			err = util.RemoveDirectory(passwordPath)
			if err != nil {
				return fmt.Errorf("failed to remove password directory: %w", err)
			}
		}
		fmt.Printf("Removed password for %s", site)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
