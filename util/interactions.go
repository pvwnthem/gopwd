package util

import (
	"fmt"
	"runtime"
	"syscall"

	"golang.org/x/term"
)

func GetPassword() (string, error) {
	fmt.Print("Enter password: ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}

	fmt.Print("Confirm password: ")
	confirm, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}

	if string(password) != string(confirm) {
		return "", fmt.Errorf("passwords do not match")
	}

	return string(password), nil
}

func ConfirmAction() (bool, error) {
	fmt.Print("Are you sure? [y/N] ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return false, err
	}

	if input == "y" || input == "Y" {
		return true, nil
	}

	return false, nil
}

func GetTextEditor() string {
	if runtime.GOOS == "windows" {
		return "notepad"
	}
	return "nano"
}
