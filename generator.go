package jpfaker

type randomSource interface {
	Intn(n int) int
}

// Generator is the entry point for all fake-data generators in this package.
type Generator struct {
	rng randomSource
}

// New creates a Generator with the provided options.
func New(opts ...Option) *Generator {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.rng == nil {
		cfg.rng = defaultConfig().rng
	}

	return &Generator{
		rng: cfg.rng,
	}
}

// Person returns a generator for Japanese personal names.
func (g *Generator) Person() PersonGenerator {
	return PersonGenerator{g: g}
}

// Address returns a generator for Japanese postal addresses.
func (g *Generator) Address() AddressGenerator {
	return AddressGenerator{g: g}
}

// Phone returns a generator for Japanese phone numbers.
func (g *Generator) Phone() PhoneGenerator {
	return PhoneGenerator{g: g}
}

// Company returns a generator for company-related values.
func (g *Generator) Company() CompanyGenerator {
	return CompanyGenerator{g: g}
}
