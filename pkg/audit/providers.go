package audit

import (
	"fmt"
	"strings"
)

var DefaultProvider *Provider = &Provider{
	Name: "default",
	Process: func(in string) (bool, string) {
		message := make([]string, 0)
		secure := true
		if len(in) <= 32 {
			message = append(message, fmt.Sprintf("password is too short (%d characters) should be at least 32", len(in)))
			secure = false
		}

		// symbols check
		symbols := 0
		for _, c := range in {
			if c >= '!' && c <= '/' {
				symbols++
			}
			if c >= ':' && c <= '@' {
				symbols++
			}
			if c >= '[' && c <= '`' {
				symbols++
			}
			if c >= '{' && c <= '~' {
				symbols++
			}
		}

		if symbols < 1 {
			message = append(message, fmt.Sprintf("password should contain at least 1 symbol, found %d", symbols))
			secure = false
		}

		// digits check
		digits := 0
		for _, c := range in {
			if c >= '0' && c <= '9' {
				digits++
			}
		}

		if digits < 1 {
			message = append(message, fmt.Sprintf("password should contain at least 1 digits, found %d", digits))
			secure = false
		}

		// uppercase check
		uppercase := 0
		for _, c := range in {
			if c >= 'A' && c <= 'Z' {
				uppercase++
			}
		}

		if uppercase < 1 {
			message = append(message, fmt.Sprintf("password should contain at least 1 uppercase letter, found %d", uppercase))
			secure = false
		}
		return secure, strings.Join(message, ", ")
	},
}
