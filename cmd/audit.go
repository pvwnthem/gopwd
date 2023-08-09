package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/audit"
	"github.com/pvwnthem/gopwd/pkg/crypto"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var auditCommand = &cobra.Command{
	Use:   "audit",
	Short: "Audit the vault for weak passwords",
	RunE: func(cmd *cobra.Command, args []string) error {
		vaultPath := Path

		// Check if the vault exists
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultExistence, err)
		}
		if !vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		// Iterate through the vault
		dir, err := util.WalkDir(vaultPath)
		if err != nil {
			return fmt.Errorf("failed to walk directory: %w", err)
		}
		GPGID, err := util.GetGPGID(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrGetGPGID, err)
		}

		GPGModule := crypto.New(GPGID, crypto.Config{})

		// array of passwords
		passwords := make(map[string]string)
		for _, d := range dir {
			splitFilePath := strings.Split(d, "/")
			if !strings.HasPrefix(splitFilePath[len(splitFilePath)-1], ".") {
				passwordPath := filepath.Join(d)
				password, err := util.ReadFile(passwordPath)
				if err != nil {
					return fmt.Errorf("failed to read password file: %w", err)
				}

				decrypted, err := GPGModule.Decrypt(password)
				if err != nil {
					return fmt.Errorf(constants.ErrPasswordDecryption, err)
				}
				// this is probably a bad way to do this
				a := strings.Split(vaultPath, "/")
				m := make(map[string]bool)

				for _, item := range a {
					m[item] = true
				}

				var diff []string
				for _, item := range splitFilePath {
					if _, ok := m[item]; !ok {
						diff = append(diff, item)
					}
				}

				passwords[strings.Join(diff[:len(diff)-1], "/")] = string(decrypted)
			}
		}

		auditor := audit.New(audit.DefaultProvider)

		for k, v := range passwords {
			secure, message := auditor.Process(v)
			if secure {
				fmt.Printf("%s: secure\n", k)
			} else {
				fmt.Printf("%s: insecure, %s\n", k, message)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(auditCommand)
}
