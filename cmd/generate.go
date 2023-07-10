package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [site] [length] [flags]",
	Short: "Generates and inserts a new password into the vault",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		site := args[0]
		length := args[1]
		vaultPath := filepath.Join(Path, Name)

		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to check vault existence: %w", err)
		}
		if !vaultExists {
			return fmt.Errorf("vault does not exist at %s", vaultPath)
		}

		// Check if the password already exists
		passwordExists, err := util.Exists(filepath.Join(vaultPath, site))
		if err != nil {
			return fmt.Errorf("failed to check password existence: %w", err)
		}
		if passwordExists {
			return fmt.Errorf("password already exists")
		}

		// Create the directory
		dirPath := filepath.Join(vaultPath, site)
		err = util.CreateDirectory(dirPath)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Generate the password
		password := util.GeneratePassword(length)

		GPGID, err := util.GetGPGID(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to get gpg-id: %w", err)
		}

		GPGModule := util.NewGPGModule(GPGID, "/usr/bin/gpg")

		encryptedPassword, err := GPGModule.Encrypt([]byte(password))
		if err != nil {
			return fmt.Errorf("failed to encrypt password: %w", err)
		}

		// Create the password file
		passwordPath := filepath.Join(dirPath, "password")
		err = util.WriteBytesToFile(passwordPath, encryptedPassword)
		if err != nil {
			return fmt.Errorf("failed to write password to file: %w", err)
		}

		fmt.Printf("Inserted password for %s at %s\n", site, dirPath)
		fmt.Printf("Generated password: %s\n", password)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
