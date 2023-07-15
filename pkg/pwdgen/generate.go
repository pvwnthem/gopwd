package pwdgen

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func (g *Generator) Generate() (string, error) {

	var password strings.Builder

	for i := 0; i < g.length; i++ {

		randomBytes := make([]byte, 1)

		_, err := rand.Read(randomBytes)
		if err != nil {
			return "", fmt.Errorf("failed to generate random bytes: %w", err)
		}

		randomIndex := int(randomBytes[0]) % len(g.characters)
		character := g.characters[randomIndex]

		password.WriteString(string(character))
	}

	return password.String(), nil
}
