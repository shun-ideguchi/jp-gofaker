package jpfaker

type randomSource interface {
	Intn(n int) int
}

type Generator struct {
	rng randomSource
}

func New(opts ...Option) *Generator {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	return &Generator{
		rng: cfg.rng,
	}
}

func (g *Generator) Person() PersonGenerator {
	return PersonGenerator{g: g}
}

func (g *Generator) Address() AddressGenerator {
	return AddressGenerator{g: g}
}

func (g *Generator) Phone() PhoneGenerator {
	return PhoneGenerator{g: g}
}

func (g *Generator) Company() CompanyGenerator {
	return CompanyGenerator{g: g}
}
