package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/crypto"
	"github.com/pvwnthem/gopwd/pkg/qr"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var (
	Line int
	Qr   bool
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

		if Qr {
			qr.Generate(string(decrypted), qr.M, os.Stdout)
			return nil
		}

		// If the clipboard flag is set, copy the password to the clipboard
		if Copy {
			err = clipboard.WriteAll(string(decrypted))
			if err != nil {
				return fmt.Errorf("failed to copy password to clipboard: %w", err)
			}
			fmt.Printf("Copied password for %s to clipboard", site)
			return nil
		}

		// If the line flag is set, print or copy only the provided line number
		lines := strings.Split(string(decrypted), "\n")
		if Line > 0 {
			line := lines[Line-1]
			if line == "" {
				return fmt.Errorf("line %d is empty", Line)
			}
			if Copy {
				err = clipboard.WriteAll(line)
				if err != nil {
					return fmt.Errorf("failed to copy line to clipboard: %w", err)
				}
				fmt.Printf("Copied line %d for %s to clipboard", Line, site)
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
	showCmd.Flags().IntVarP(&Line, "line", "l", 0, "print or copy only the provided line number")
	showCmd.Flags().BoolVarP(&Qr, "qr", "q", false, "print the password as a QR code")
	rootCmd.AddCommand(showCmd)
}
