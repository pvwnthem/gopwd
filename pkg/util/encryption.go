package util

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetGPGID(path string) (string, error) {
	file, err := os.Open(filepath.Join(path, ".gpg-id"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	gpgIDBytes := make([]byte, 0)
	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			gpgIDBytes = append(gpgIDBytes, buffer[:n]...)
		}
		if err != nil {
			break
		}
	}

	return strings.TrimSpace(string(gpgIDBytes)), nil
}

func GeneratePassword(length string) (string, error) {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+[]{};':\",./<>?"

	// Convert the length string to an integer
	n, err := strconv.Atoi(length)
	if err != nil {
		return "", fmt.Errorf("invalid length: %w", err)
	}

	// Generate the password
	var password strings.Builder
	for i := 0; i < n; i++ {
		randomBytes := make([]byte, 1)
		_, err := rand.Read(randomBytes)
		if err != nil {
			return "", fmt.Errorf("failed to generate random bytes: %w", err)
		}
		randomIndex := int(randomBytes[0]) % len(characters)
		character := characters[randomIndex]
		password.WriteByte(character)
	}

	return password.String(), nil
}
