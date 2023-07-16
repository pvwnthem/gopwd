package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/crypto"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert [site] [flags]",
	Short: "Inserts a new password into the vault",
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
			return fmt.Errorf(constants.ErrVaultDoesExist, vaultPath)
		}

		// Check if the password already exists
		passwordExists, err := util.Exists(filepath.Join(vaultPath, site, "password"))
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordExistence, err)
		}
		if passwordExists {
			return fmt.Errorf(constants.ErrPasswordDoesExist)
		}

		password, err := util.GetPassword()
		if err != nil {
			return fmt.Errorf("failed to get password: %w", err)
		}

		// Create the directory
		dirPath := filepath.Join(vaultPath, site)
		err = util.CreateDirectory(dirPath)
		if err != nil {
			return fmt.Errorf(constants.ErrDirectoryCreation, err)
		}

		// Create the password file
		passwordPath := filepath.Join(dirPath, "password")
		err = util.CreateFile(passwordPath)
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf("failed to create password file: %w", err)
		}

		GPGID, err := util.GetGPGID(vaultPath)
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf(constants.ErrGetGPGID, err)
		}

		GPGModule := crypto.New(GPGID, crypto.Config{})

		encryptedPassword, err := GPGModule.Encrypt([]byte(password))
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf(constants.ErrPasswordEncryption, err)
		}

		err = util.WriteBytesToFile(passwordPath, encryptedPassword)
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf(constants.ErrPasswordWrite, err)
		}

		fmt.Printf("Inserted password for %s at %s", site, dirPath)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}
