package audit

type Provider struct {
	Name    string
	Process func(string) bool
}
