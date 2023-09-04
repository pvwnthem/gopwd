package pwdgen

import "github.com/torbenconto/zeus"

type Generator struct {
	length     int
	characters string
}

var (
	Digits       = "0123456789"
	Upper        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lower        = "abcdefghijklmnopqrstuvwxyz"
	Syms         = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	CharAlpha    = zeus.Concat(Upper, Lower)
	CharAlphaNum = zeus.Concat(Digits, Upper, Lower)
	CharAll      = zeus.Concat(Digits, Upper, Lower, Syms)
)

func New(length int, characters string) *Generator {
	return &Generator{
		length:     length,
		characters: characters,
	}
}
