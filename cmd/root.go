package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pvwnthem/gopwd/cmd/config"
	"github.com/pvwnthem/gopwd/cmd/vault"
	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/util"
	"github.com/spf13/cobra"
)

var configFile string

var (
	Path    string
	Version string
	Copy    bool
)

var rootCmd = &cobra.Command{
	Use:     "gopwd",
	Short:   "A cli password manager written in go",
	Long:    "gopwd is an encrypted cli password manager (similar to password-store) written in golang",
	Version: Version,

	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the vault exists
		vaultPath := Path
		pathExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultExistence, err)
		}
		if !pathExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		isVault, err := util.Exists(filepath.Join(vaultPath, ".vault"))
		if err != nil {
			return fmt.Errorf("failed to check if the provided path is a vault %w", err)
		}
		if !isVault {
			return fmt.Errorf("provided path is not a vault, does it have a .vault file?")
		}

		fmt.Println(strings.Split(Path, "/")[len(strings.Split(Path, "/"))-1]) // print name of vault on top of dir structure
		err = util.PrintDirectoryTree(Path, "")
		if err != nil {
			return fmt.Errorf("error printing directory tree: %w", err)
		}

		return nil

	},
}

func Execute() {
	Path, configFile, _ = util.InitConfig(Path, configFile)

	rootCmd.Version = Version
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(vault.VaultCmd)
	rootCmd.AddCommand(config.ConfigCmd)
}

func init() {
	addSubcommandPalettes()
	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "The path of the vault")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Config file (default is $HOME/.gopwd/config.json)")
	rootCmd.PersistentFlags().BoolVarP(&Copy, "copy", "c", false, "copy the password to the clipboard and don't print it to stdout")
}
