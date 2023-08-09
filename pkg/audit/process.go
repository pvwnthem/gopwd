package audit

func (a *Audit) Process(in string) bool {
	return a.provider.Process(in)
}
