package util

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func GetPassword() (string, error) {
	fmt.Print("Enter password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}

	fmt.Print("Confirm password: ")
	confirm, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		return "", err
	}

	if string(password) != string(confirm) {
		return "", fmt.Errorf("passwords do not match")
	}

	return string(password), nil
}
