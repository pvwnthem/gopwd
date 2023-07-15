package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pvwnthem/gopwd/constants"
	"github.com/pvwnthem/gopwd/pkg/util"
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

		configExists, err := util.Exists(filepath.Join(util.GetHomeDir(), ".gopwd", "config.json"))
		if err != nil {
			return fmt.Errorf(constants.ErrConfigExistence, err)
		}
		if configExists {
			return errors.New(constants.ErrConfigDoesExist)
		}
		// Read the config file
		configJSON, err := os.ReadFile(configPath)
		if err != nil {
			return fmt.Errorf(constants.ErrConfigRead, err)
		}

		// Unmarshal the config object from JSON
		var config util.Config
		err = json.Unmarshal(configJSON, &config)
		if err != nil {
			return fmt.Errorf(constants.ErrJSONUnmarshal, err)
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
			return fmt.Errorf(constants.ErrJSONMarshal, err)
		}

		// Write the JSON to the config file
		err = os.WriteFile(configPath, configJSON, 0644)
		if err != nil {
			return fmt.Errorf(constants.ErrConfigWrite, err)
		}

		fmt.Println("Successfully modified config")
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(setCmd)
}
