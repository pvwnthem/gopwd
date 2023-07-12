package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pvwnthem/gopwd/cmd/config"
	"github.com/pvwnthem/gopwd/cmd/vault"
	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var configFile string

var (
	Path string
	Name string
)

var rootCmd = &cobra.Command{
	Use:   "gopwd",
	Short: "A cli password manager written in go",
	Long:  "gopwd is an encrypted cli password manager (similar to password-store) written in golang",

	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if the vault exists
		vaultPath := filepath.Join(Path, Name)
		vaultExists, err := util.Exists(vaultPath)
		if err != nil {
			return fmt.Errorf(constants.ErrVaultExistence, err)
		}
		if !vaultExists {
			return fmt.Errorf(constants.ErrVaultDoesNotExist, vaultPath)
		}

		fmt.Println(Name) // print name of vault on top of dir structure
		err = util.PrintDirectoryTree(filepath.Join(Path, Name), "")
		if err != nil {
			return fmt.Errorf("error printing directory tree: %w", err)
		}

		return nil

	},
}

func Execute() {
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
	cobra.OnInitialize(func() { Path, Name, configFile, _ = util.InitConfig(Path, Name, configFile) })

	addSubcommandPalettes()
	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "The path of the vault")
	rootCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "The name of the vault")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Config file (default is $HOME/.gopwd/config.json)")
}
