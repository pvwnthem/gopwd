package util

import (
	"fmt"
	"io/ioutil"
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
	tmpfile, err := ioutil.TempFile("", "gpg-encrypt-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpfile.Name())

	err = os.WriteFile(tmpfile.Name(), plaintext, 0600)
	if err != nil {
		return nil, err
	}

	outputFile := fmt.Sprintf("%s.gpg", tmpfile.Name())
	defer os.Remove(outputFile)

	cmd := exec.Command(g.GPGPath, "--encrypt", "--armor", "--recipient", g.GPGID, "--output", outputFile, tmpfile.Name())
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return os.ReadFile(outputFile)
}

func (g *GPGModule) Decrypt(ciphertext []byte) ([]byte, error) {
	tmpfile, err := ioutil.TempFile("", "gpg-decrypt-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpfile.Name())

	err = os.WriteFile(tmpfile.Name(), ciphertext, 0600)
	if err != nil {
		return nil, err
	}

	outputFile := fmt.Sprintf("%s.decrypted", tmpfile.Name())
	defer os.Remove(outputFile)

	cmd := exec.Command(g.GPGPath, "--decrypt", "--output", outputFile, tmpfile.Name())
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return os.ReadFile(outputFile)
}

func GetGPGID(path string) (string, error) {
	gpgID, err := os.ReadFile(filepath.Join(path, ".gpg-id"))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(gpgID)), nil
}
