package audit

type Provider struct {
	Name string
	//                 is secure, message
	Process func(string) (bool, []string)
}
