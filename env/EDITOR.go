package env

import (
	"os"
	"runtime"
)

func EDITOR() string {
	if os.Getenv("EDITOR") != "" {
		return os.Getenv("EDITOR")
	}

	if runtime.GOOS == "windows" {
		return "notepad"
	}

	if runtime.GOOS == "darwin" {
		return "nano"
	}

	return "vi"
}
