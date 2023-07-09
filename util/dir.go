package util

import "os"

func GetHomeDir() string {
	homeDir, homeDirErr := os.UserHomeDir()

	if homeDirErr != nil {
		panic(homeDirErr)
	}

	return homeDir
}
