package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	DefaultPath = filepath.Join(GetHomeDir(), ".gopwd")
	DefaultName = "vault"
)

// Config holds the configuration data
type Config struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

func LoadConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func InitConfig(Path string, Name string, configFile string) (string, string, string, error) {
	// Use default configuration file path
	if configFile == "" {
		configFile = filepath.Join(GetHomeDir(), ".gopwd", "config.json")
	}

	// Check if the config file exists
	_, err := os.Stat(configFile)
	if err != nil {
		// If the config file doesn't exist, use default values
		Path = DefaultPath
		Name = DefaultName
		return Path, Name, configFile, nil
	}

	// Load configuration from file
	cfg, err := LoadConfig(configFile)
	if err != nil {
		fmt.Println("Failed to load config file:", err)
		os.Exit(1)
		return "", "", "", err
	}

	// Override flags with configuration values
	if Path == "" {
		Path = cfg.Path
	}
	if Name == "" {
		Name = cfg.Name
	}

	return Path, Name, configFile, nil

}
