package audit

import (
	"fmt"
	"unicode"

	"github.com/pvwnthem/gopwd/pkg/hibp"
)

var DefaultProvider *Provider = &Provider{
	Name: "default",
	Process: func(in string) (bool, []string, error) {

		var (
			secure  bool = true
			message []string

			symbols int
			digits  int
			upper   int
		)

		if len(in) < 32 {
			message = append(message, fmt.Sprintf("Password is too short (%d characters), should be at least 32", len(in)))
			secure = false
		}

		for _, c := range in {
			switch {
			case unicode.IsSymbol(c) || unicode.IsPunct(c):
				symbols++
			case unicode.IsDigit(c):
				digits++
			case unicode.IsUpper(c):
				upper++
			}
		}

		if symbols < 1 {
			message = append(message, "Password should contain at least one symbol")
			secure = false
		}

		if digits < 1 {
			message = append(message, "Password should contain at least one digit")
			secure = false
		}

		if upper < 1 {
			message = append(message, "Password should contain at least one upper case letter")
			secure = false
		}

		return secure, message, nil
	},
}

var HibpProvider *Provider = &Provider{
	Name: "hibp (haveibeenpwned.com)",
	Process: func(in string) (bool, []string, error) {
		check, err := hibp.Check(in)
		if err != nil {
			return false, nil, err
		}

		if check {
			return false, []string{"Password has been compromised"}, nil
		} else {
			return true, []string{""}, nil
		}
	},
}
