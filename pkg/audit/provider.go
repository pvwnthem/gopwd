package audit

type Provider struct {
	Name string
	//                 is secure, message
	//                     |        |
	//                     v		v
	Process func(string) (bool, []string, error)
}
