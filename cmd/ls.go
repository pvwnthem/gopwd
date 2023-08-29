package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls [password] [flags]",
	Short: "Copies a password from one site to another",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		password := args[0]

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
		passwordPath := filepath.Join(vaultPath, password)
		passwordExists, err := util.Exists(passwordPath)
		if err != nil {
			return fmt.Errorf(constants.ErrPasswordExistence, err)
		}
		if !passwordExists {
			return fmt.Errorf(constants.ErrPasswordDoesNotExist)
		}

		isPassword, _ := util.Exists(filepath.Join(passwordPath, "password"))

		if isPassword {
			fmt.Printf("%s (password) \n", password)
		} else {
			fmt.Printf("%s \n", password)
		}

		util.PrintDirectoryTree(passwordPath, "")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
