package cmd

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/crypto"
	"github.com/pvwnthem/gopwd/pkg/pwdgen"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var (
	Memorable bool
	Symbols   bool
)

var generateCmd = &cobra.Command{
	Use:   "generate [site] [length] [flags]",
	Short: "Generates and inserts a new password into the vault",
	Args:  cobra.ExactArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		site := args[0]
		length := args[1]
		vaultPath := Path

		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, err)
		}
		if !vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		// Check if the password already exists
		passwordExists, err := util.Exists(filepath.Join(vaultPath, site, "password"))
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
		len, err := strconv.Atoi(length)
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf("failed to convert length to int: %w", err)
		}
		generator := pwdgen.NewGenerator(len, pwdgen.CharAll)
		var password string
		if Memorable {
			password, err = generator.GenerateMemorable(true, !Symbols)
		} else {
			password, err = generator.Generate()
		}
		if err != nil {
			util.RemoveDirectory(dirPath)
			return fmt.Errorf("failed to generate password %w", err)
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
	generateCmd.Flags().BoolVarP(&Memorable, "memorable", "m", false, "Generate a memorable password")
	generateCmd.Flags().BoolVar(&Symbols, "no-symbols", false, "Generate a password with no symbols")
	rootCmd.AddCommand(generateCmd)
}
