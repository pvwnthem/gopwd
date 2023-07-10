package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
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

	data, err := ioutil.ReadAll(file)
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
