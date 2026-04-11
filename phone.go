package jpfaker

import "fmt"

type PhoneGenerator struct {
	g *Generator
}

func (p PhoneGenerator) Mobile() string {
	prefixes := []string{"070", "080", "090"}

	return fmt.Sprintf(
		"%s-%04d-%04d",
		prefixes[p.g.rng.Intn(len(prefixes))],
		p.g.rng.Intn(10000),
		p.g.rng.Intn(10000),
	)
}

func (p PhoneGenerator) Landline() string {
	areaCodes := []string{"03", "06", "011", "052", "075", "092"}
	areaCode := areaCodes[p.g.rng.Intn(len(areaCodes))]

	if len(areaCode) == 2 {
		return fmt.Sprintf("%s-%04d-%04d", areaCode, p.g.rng.Intn(10000), p.g.rng.Intn(10000))
	}

	return fmt.Sprintf("%s-%03d-%04d", areaCode, p.g.rng.Intn(1000), p.g.rng.Intn(10000))
}

func (p PhoneGenerator) TollFree() string {
	prefixes := []string{"0120", "0800"}

	return fmt.Sprintf(
		"%s-%03d-%03d",
		prefixes[p.g.rng.Intn(len(prefixes))],
		p.g.rng.Intn(1000),
		p.g.rng.Intn(1000),
	)
}
