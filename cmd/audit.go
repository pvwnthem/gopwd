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

				passwords[d] = string(decrypted)
			}
		}

		auditor := audit.New(&audit.Provider{})

		for k, v := range passwords {
			fmt.Printf("%s %s\n", k, auditor.Process(v))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(auditCommand)
}
