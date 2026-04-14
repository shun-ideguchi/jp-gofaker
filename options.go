package jpfaker

import (
	"math/rand"
	"time"
)

type config struct {
	rng randomSource
}

// Option configures a Generator.
type Option func(*config)

func defaultConfig() config {
	return config{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// WithSeed makes generated values deterministic for the same seed.
func WithSeed(seed int64) Option {
	return func(cfg *config) {
		cfg.rng = rand.New(rand.NewSource(seed))
	}
}

func withRandomSource(rng randomSource) Option {
	return func(cfg *config) {
		cfg.rng = rng
	}
}
