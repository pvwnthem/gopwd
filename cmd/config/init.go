package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a config file at the specified path",
	Long:  "",

	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := filepath.Join(util.GetHomeDir(), ".gopwd", "config.json")

		// Check if the config file exists
		configExists, err := util.Exists(filepath.Join(util.GetHomeDir(), ".gopwd", "config.json"))
		if err != nil {
			return fmt.Errorf(constants.ErrConfigExistence, err)
		}
		if configExists {
			return errors.New(constants.ErrConfigDoesExist)
		}

		// Create the config file
		err = util.CreateFile(configPath)
		if err != nil {
			return fmt.Errorf("failed to create config file: %w", err)
		}

		// Create the config object
		config := util.Config{
			Path: Path,
			Name: Name,
		}

		// Marshal the config object to JSON
		configJSON, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf(constants.ErrJSONMarshal, err)
		}

		// Write the JSON to the config file
		err = os.WriteFile(configPath, configJSON, 0644)
		if err != nil {
			return fmt.Errorf(constants.ErrConfigWrite, err)
		}
		return nil
	},
}

func init() {
	initCmd.MarkPersistentFlagRequired("name")
	initCmd.MarkPersistentFlagRequired("path")
	ConfigCmd.AddCommand(initCmd)
}
