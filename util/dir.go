package util

import "os"

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	return homeDir
}
