package jpfaker

import (
	"strings"

	"github.com/shun-ideguchi/jp-gofaker/internal/dataset"
)

// CompanyGenerator generates company-related values.
type CompanyGenerator struct {
	g *Generator
}

// Name returns a generated company name.
func (c CompanyGenerator) Name() string {
	legalType, _ := pickOne(c.g.rng, dataset.CompanyLegalTypes)
	prefix, _ := pickOne(c.g.rng, dataset.CompanyPrefixes)
	core, _ := pickOne(c.g.rng, dataset.CompanyCores)

	return strings.TrimSpace(legalType + prefix + core)
}

// Department returns a generated department name.
func (c CompanyGenerator) Department() string {
	department, _ := pickOne(c.g.rng, dataset.CompanyDepartments)
	return department
}

// Title returns a generated job title.
func (c CompanyGenerator) Title() string {
	title, _ := pickOne(c.g.rng, dataset.CompanyTitles)
	return title
}
