package crypt

var (
	defaultArgs = []string{"--quiet", "--yes", "--compress-algo=none", "--no-encrypt-to", "--no-auto-check-trustdb"}
)

type GPG struct {
	id         string
	binaryPath string
	args       []string
}

type Config struct {
	BinaryPath string
	Args       []string
}

func New(id string, config Config) *GPG {
	gpg := &GPG{
		id: id,
	}

	// Set default values
	if config.BinaryPath == "" {
		config.BinaryPath = "gpg"
	}
	if len(config.Args) == 0 {
		config.Args = defaultArgs
	}

	// Set config values
	gpg.binaryPath = config.BinaryPath
	gpg.args = config.Args

	return gpg
}

func (g *GPG) Binary() string {
	return g.binaryPath
}

func (g *GPG) ID() string {
	return g.id
}

func (g *GPG) Args() []string {
	return g.args
}
