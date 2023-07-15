package pwdgen

import (
	"strings"

	"github.com/pvwnthem/gopwd/pkg/crypto"
)

func (g *Generator) GenerateMemorable(capitals bool, symbols bool) (string, error) {

	var password strings.Builder

	for password.Len() < g.length {
		if capitals && crypto.RandomInt(2) == 1 {
			password.WriteString(strings.ToTitle(wordlist[crypto.RandomInt(len(wordlist))]))
		} else {
			password.WriteString(wordlist[crypto.RandomInt(len(wordlist))])
		}

		password.WriteByte(Digits[crypto.RandomInt(len(Digits))])

		if !symbols {
			continue
		}

		password.WriteByte(Syms[crypto.RandomInt(len(Syms))])
	}

	return password.String(), nil
}
