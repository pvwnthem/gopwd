package crypto

import (
	"bytes"
	"os"
	"os/exec"
)

func (g *GPG) Decrypt(ciphertext []byte) ([]byte, error) {
	args := append(g.Args(), "--decrypt")

	buffer := &bytes.Buffer{}

	cmd := exec.Command(g.Binary(), args...)
	cmd.Stdin = bytes.NewReader(ciphertext)
	cmd.Stdout = buffer
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	return buffer.Bytes(), err
}
