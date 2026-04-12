package jpfaker

import "github.com/shun-ideguchi/jp-gofaker/internal/dataset"

type CompanyGenerator struct {
	g *Generator
}

func (c CompanyGenerator) Name() string {
	legalType := dataset.CompanyLegalTypes[c.g.rng.Intn(len(dataset.CompanyLegalTypes))]
	prefix := dataset.CompanyPrefixes[c.g.rng.Intn(len(dataset.CompanyPrefixes))]
	core := dataset.CompanyCores[c.g.rng.Intn(len(dataset.CompanyCores))]

	return legalType + prefix + core
}

func (c CompanyGenerator) Department() string {
	return dataset.CompanyDepartments[c.g.rng.Intn(len(dataset.CompanyDepartments))]
}

func (c CompanyGenerator) Title() string {
	return dataset.CompanyTitles[c.g.rng.Intn(len(dataset.CompanyTitles))]
}
