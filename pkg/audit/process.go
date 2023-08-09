package audit

func (a *Audit) Process(in string) (bool, string) {
	return a.provider.Process(in)
}
