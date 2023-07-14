package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
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
			return fmt.Errorf(constants.ErrVaultDoesNotExist, err)
		}
		if !vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		// Check if the password already exists
		passwordExists, err := util.Exists(filepath.Join(vaultPath, site))
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordExistence, err)
		}
		if passwordExists {
			return fmt.Errorf(constants.ErrPasswordDoesExist)
		}

		// Create the directory
		dirPath := filepath.Join(vaultPath, site)
		err = util.CreateDirectory(dirPath)
		if err != nil {
			return fmt.Errorf(constants.ErrDirectoryCreation, err)
		}

		// Generate the password
		password, err := util.GeneratePassword(length)
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf("failed to generate password %w", err)
		}

		GPGID, err := util.GetGPGID(vaultPath)
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf(constants.ErrGetGPGID, err)
		}

		GPGModule := util.NewGPGModule(GPGID, "/usr/bin/gpg")

		encryptedPassword, err := GPGModule.Encrypt([]byte(password))
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf(constants.ErrPasswordEncryption, err)
		}

		// Create the password file
		passwordPath := filepath.Join(dirPath, "password")
		err = util.WriteBytesToFile(passwordPath, encryptedPassword)
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf(constants.ErrPasswordWrite, err)
		}

		fmt.Printf("Inserted password for %s at %s\n", site, dirPath)
		fmt.Printf("Generated password: %s\n", password)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
