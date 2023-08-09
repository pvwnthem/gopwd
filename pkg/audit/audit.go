package audit

type Audit struct {
	provider Provider
}

func New(provider *Provider) *Audit {
	return &Audit{
		provider: *provider,
	}
}
