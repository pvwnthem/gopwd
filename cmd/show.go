package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [site]",
	Short: "Shows a password from the vault",
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
		passwordPath := filepath.Join(vaultPath, site, "password")
		passwordExists, err := util.Exists(passwordPath)
		if err != nil {
			return fmt.Errorf("failed to check password existence: %w", err)
		}
		if !passwordExists {
			return fmt.Errorf("password does not exist")
		}

		// Get the password and decrypt it
		file, err := util.ReadFile(passwordPath)
		if err != nil {
			return fmt.Errorf("failed to read password file: %w", err)
		}

		GPGID, err := util.GetGPGID(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to get gpg-id: %w", err)
		}

		GPGModule := util.NewGPGModule(GPGID, "/usr/bin/gpg")

		decrypted, err := GPGModule.Decrypt(file)
		if err != nil {
			return fmt.Errorf("failed to decrypt password: %w", err)
		}

		fmt.Printf("%s", decrypted)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
