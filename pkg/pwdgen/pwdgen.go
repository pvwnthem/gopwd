package pwdgen

type Generator struct {
	length     int
	characters string
}

const (
	Digits       = "0123456789"
	Upper        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lower        = "abcdefghijklmnopqrstuvwxyz"
	Syms         = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	CharAlpha    = Upper + Lower
	CharAlphaNum = Digits + Upper + Lower
	CharAll      = Digits + Upper + Lower + Syms
)

func New(length int, characters string) *Generator {
	return &Generator{
		length:     length,
		characters: characters,
	}
}
