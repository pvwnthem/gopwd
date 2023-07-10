package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pvwnthem/gopwd/cmd/vault"
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

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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
}

func init() {
	cobra.OnInitialize(initConfig)

	addSubcommandPalettes()
	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "The path of the vault")
	rootCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "The name of the vault")
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "Config file (default is $HOME/.gopwd/config.json)")
}

func initConfig() {
	// Use default configuration file path
	configFile = filepath.Join(util.GetHomeDir(), ".gopwd", "config.json")

	// Check if the config file exists
	_, err := os.Stat(configFile)
	if err != nil {
		// If the config file doesn't exist, use default values
		Path = filepath.Join(util.GetHomeDir(), ".gopwd")
		Name = "vault"
		fmt.Println("Using default values.")
		return
	}

	// Load configuration from file
	cfg, err := util.LoadConfig(configFile)
	if err != nil {
		fmt.Println("Failed to load config file:", err)
		os.Exit(1)
	}

	// Override flags with configuration values
	if Path == "" {
		Path = cfg.Path
	}
	if Name == "" {
		Name = cfg.Name
	}

	// Print success message
	fmt.Println("Configuration loaded successfully!")
}
