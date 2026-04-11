package jpfaker

import (
	"math/rand"
	"time"
)

type config struct {
	rng randomSource
}

type Option func(*config)

func defaultConfig() config {
	return config{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func WithSeed(seed int64) Option {
	return func(cfg *config) {
		cfg.rng = rand.New(rand.NewSource(seed))
	}
}
