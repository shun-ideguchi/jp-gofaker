package jpfaker

import "github.com/shun-ideguchi/jp-gofaker/internal/dataset"

// CompanyGenerator generates company-related values.
type CompanyGenerator struct {
	g *Generator
}

// Name returns a generated company name.
func (c CompanyGenerator) Name() string {
	legalType := dataset.CompanyLegalTypes[c.g.rng.Intn(len(dataset.CompanyLegalTypes))]
	prefix := dataset.CompanyPrefixes[c.g.rng.Intn(len(dataset.CompanyPrefixes))]
	core := dataset.CompanyCores[c.g.rng.Intn(len(dataset.CompanyCores))]

	return legalType + prefix + core
}

// Department returns a generated department name.
func (c CompanyGenerator) Department() string {
	return dataset.CompanyDepartments[c.g.rng.Intn(len(dataset.CompanyDepartments))]
}

// Title returns a generated job title.
func (c CompanyGenerator) Title() string {
	return dataset.CompanyTitles[c.g.rng.Intn(len(dataset.CompanyTitles))]
}
