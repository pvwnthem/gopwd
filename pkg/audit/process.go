package audit

func (a *Audit) Process(in string) (bool, []string, error) {
	return a.provider.Process(in)
}
