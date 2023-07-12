package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [site] [flags]",
	Short: "Opens the password file for editing",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		site := args[0]
		vaultPath := filepath.Join(Path, Name)

		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultExistence, err)
		}
		if !vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		// Check if the password exists
		passwordExists, err := util.Exists(filepath.Join(vaultPath, site))
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordExistence, err)
		}
		if !passwordExists {
			return fmt.Errorf(constants.ErrPasswordDoesNotExist)
		}

		GPGID, err := util.GetGPGID(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrGetGPGID, err)
		}

		GPGModule := util.NewGPGModule(GPGID, util.GetGPGPath())
		// Decrypt the password file
		passwordPath := filepath.Join(vaultPath, site, "password")
		file, _ := util.ReadBytesFromFile(passwordPath)
		decryptedContent, err := GPGModule.Decrypt(file)
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordDecryption, err)
		}

		// Create a temporary file with the decrypted content
		tmpfile := util.CreateTempFileFromBytes(decryptedContent)

		// Open the temporary file using Nano
		cmde := exec.Command(util.GetTextEditor(), tmpfile.Name())
		cmde.Stdin = os.Stdin
		cmde.Stdout = os.Stdout
		cmde.Stderr = os.Stderr

		err = cmde.Run()
		if err != nil {
			return fmt.Errorf("failed to open password file for editing: %w", err)
		}

		// Read the edited content from the temporary file
		editedContent, err := util.ReadBytesFromFile(tmpfile.Name())
		if err != nil {
			return fmt.Errorf("failed to read edited content: %w", err)
		}

		// Re-encrypt and write the edited content to the password file
		encryptedContent, err := GPGModule.Encrypt(editedContent)
		if err != nil {
			return fmt.Errorf("failed to encrypt edited content: %w", err)
		}

		err = util.WriteBytesToFile(passwordPath, encryptedContent)
		if err != nil {
			return fmt.Errorf("failed to re-encrypt and write to password file: %w", err)
		}

		// Remove the temporary file
		err = os.Remove(tmpfile.Name())
		if err != nil {
			return fmt.Errorf("failed to remove temporary file: %w", err)
		}

		fmt.Printf("Finished editing password for %s\n", site)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
