package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/crypto"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [site]",
	Short: "Shows a password from the vault",
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
		passwordPath := filepath.Join(vaultPath, site, "password")
		passwordExists, err := util.Exists(passwordPath)
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordExistence, err)
		}
		if !passwordExists {
			return fmt.Errorf(constants.ErrPasswordDoesNotExist)
		}

		// Get the password and decrypt it
		file, err := util.ReadFile(passwordPath)
		if err != nil {
			return fmt.Errorf("failed to read password file: %w", err)
		}

		GPGID, err := util.GetGPGID(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrGetGPGID, err)
		}

		GPGModule := crypto.New(GPGID, crypto.Config{})

		decrypted, err := GPGModule.Decrypt(file)
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordDecryption, err)
		}

		// If the clipboard flag is set, copy the password to the clipboard
		clipboardFlag, _ := cmd.Flags().GetBool("copy")
		if clipboardFlag {
			err = clipboard.WriteAll(string(decrypted))
			if err != nil {
				return fmt.Errorf("failed to copy password to clipboard: %w", err)
			}
			fmt.Printf("Copied password for %s to clipboard", site)
			return nil
		}

		// If the line flag is set, print or copy only the provided line number
		lineNumber, _ := cmd.Flags().GetInt("line")
		lines := strings.Split(string(decrypted), "\n")
		if lineNumber > 0 {
			line := lines[lineNumber-1]
			if line == "" {
				return fmt.Errorf("line %d is empty", lineNumber)
			}
			if clipboardFlag {
				err = clipboard.WriteAll(line)
				if err != nil {
					return fmt.Errorf("failed to copy line to clipboard: %w", err)
				}
				fmt.Printf("Copied line %d for %s to clipboard", lineNumber, site)
			} else {
				fmt.Printf("%s", string(line))
			}
			return nil
		}

		fmt.Printf("%s \n", decrypted)

		return nil
	},
}

func init() {
	showCmd.Flags().BoolP("copy", "c", false, "copy the password to the clipboard and don't print it to stdout")
	showCmd.Flags().IntP("line", "l", 0, "print or copy only the provided line number")
	rootCmd.AddCommand(showCmd)
}
