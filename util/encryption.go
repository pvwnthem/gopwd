package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type GPGModule struct {
	GPGID   string // GPG ID used for encryption and decryption
	GPGPath string // Path to the GnuPG executable
}

func NewGPGModule(gpgID, gpgPath string) *GPGModule {
	return &GPGModule{
		GPGID:   gpgID,
		GPGPath: gpgPath,
	}
}

func (g *GPGModule) Encrypt(plaintext []byte) ([]byte, error) {
	tmpfile, err := createTempFile("", "gpg-encrypt-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpfile.Name())

	err = writeFile(tmpfile.Name(), plaintext, 0600)
	if err != nil {
		return nil, err
	}

	outputFile := fmt.Sprintf("%s.gpg", tmpfile.Name())
	defer os.Remove(outputFile)

	cmd := exec.Command(g.GPGPath, "--encrypt", "--armor", "--recipient", g.GPGID, "--output", outputFile, tmpfile.Name())
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return readFile(outputFile)
}

func (g *GPGModule) Decrypt(ciphertext []byte) ([]byte, error) {
	tmpfile, err := createTempFile("", "gpg-decrypt-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpfile.Name())

	err = writeFile(tmpfile.Name(), ciphertext, 0600)
	if err != nil {
		return nil, err
	}

	outputFile := fmt.Sprintf("%s.decrypted", tmpfile.Name())
	defer os.Remove(outputFile)

	cmd := exec.Command(g.GPGPath, "--decrypt", "--output", outputFile, tmpfile.Name())
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return readFile(outputFile)
}

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

func GeneratePassword(length string) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+[]{};':\",./<>?"

	// Convert the length string to an integer
	n, _ := new(big.Int).SetString(length, 10)
	passwordLength := int(n.Int64())

	// Generate the password
	var password []byte
	for i := 0; i < passwordLength; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		character := characters[randomIndex.Int64()]
		password = append(password, byte(character))
	}

	return string(password)
}
