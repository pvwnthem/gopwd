package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/audit"
	"github.com/pvwnthem/gopwd/pkg/colors"
	"github.com/pvwnthem/gopwd/pkg/crypto"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var (
	Hibp   bool
	Custom bool

	min_length  int
	min_digits  int
	min_symbols int
	min_upper   int
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

		var provider *audit.Provider

		auditor := audit.New(provider)

		// check for duplicates, returns the names of the duplicate passwords and a bool indicating whether or not there are any duplicates
		/*
			it should be noted that this function will not detect duplicate passwords that are with other data in their file
			this could be fix by iterating throught file lines as well but that would still not be 100% accurate
			checking with a deep search might return emails or usernames given as duplicate passwords and there is no way to distinguish these from passwords.
		*/
		output1, output2, duplicate := audit.CheckDuplicates(passwords)

		if duplicate {
			fmt.Println("Duplicate passwords found: ")
			fmt.Println(colors.Redf(output1))
			fmt.Println(colors.Redf(output2))
			fmt.Println("------------------------")
		}

		for k, v := range passwords {
			secure, messages, err := auditor.Process(v)

			if err != nil {
				return fmt.Errorf("error processing audit provider %s : %v", provider.Name, err)
			}
			if secure {
				fmt.Print(colors.Greenf("%s: secure\n", k))
			} else {
				fmt.Print(colors.Redf("%s: insecure\n", k))
				for _, message := range messages {
					fmt.Printf(">	%s\n", message)
				}
			}
		}

		return nil
	},
}

func init() {
	// only one provider at a time is supported currently
	auditCommand.Flags().BoolVar(&Hibp, "hibp", false, "Check passwords against haveibeenpwned.com")
	auditCommand.Flags().BoolVar(&Custom, "custom", false, "Use a custom ruleset for auditing passwords")

	auditCommand.Flags().IntVarP(&min_length, "length", "l", 32, "Minimum allowed length of the password (default is 32)")
	auditCommand.Flags().IntVarP(&min_digits, "digits", "d", 1, "Minimum amount of digits allowed in the password (default is 1)")
	auditCommand.Flags().IntVarP(&min_symbols, "symbols", "s", 1, "Miniumum amount of symbols allowed in the password (default is 1)")
	auditCommand.Flags().IntVarP(&min_upper, "upper", "u", 1, "Miniumum amount of uppercase characters allowed in the password (default is 1)")

	rootCmd.AddCommand(auditCommand)
}
