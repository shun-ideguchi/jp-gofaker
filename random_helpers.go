package jpfaker

func pickOne[T any](rng randomSource, values []T) (T, bool) {
	var zero T

	if rng == nil || len(values) == 0 {
		return zero, false
	}

	return values[rng.Intn(len(values))], true
}
