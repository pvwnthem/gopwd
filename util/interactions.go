package util

import (
	"errors"
	"fmt"
)

func GetPassword() (string, error) {
	var password string
	var confirm string

	fmt.Print("Enter password: ")
	_, err := fmt.Scan(&password)
	if err != nil {
		return "", errors.New("failed to read password")
	}

	fmt.Print("Confirm password: ")
	_, err = fmt.Scan(&confirm)
	if err != nil {
		return "", errors.New("failed to read confirmation")
	}

	if password != confirm {
		return "", errors.New("passwords do not match")
	}

	return password, nil
}
