package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pvwnthem/gopwd/util"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [field] [value]",
	Short: "Modify config fields",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		field := args[0]
		value := args[1]

		if field != "path" && field != "name" {
			return errors.New("invalid field name, should be either 'path' or 'name'")
		}

		configPath := filepath.Join(util.GetHomeDir(), ".gopwd", "config.json")

		// Read the config file
		configJSON, err := os.ReadFile(configPath)
		if err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}

		// Unmarshal the config object from JSON
		var config util.Config
		err = json.Unmarshal(configJSON, &config)
		if err != nil {
			return fmt.Errorf("failed to unmarshal config from JSON: %w", err)
		}

		// Modify the specified field
		switch field {
		case "path":
			config.Path = value
		case "name":
			config.Name = value
		}

		// Marshal the modified config object to JSON
		configJSON, err = json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal config to JSON: %w", err)
		}

		// Write the JSON to the config file
		err = os.WriteFile(configPath, configJSON, 0644)
		if err != nil {
			return fmt.Errorf("failed to write config file: %w", err)
		}

		fmt.Println("Successfully modified config")
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(setCmd)
}
