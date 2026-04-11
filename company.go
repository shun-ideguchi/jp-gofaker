package jpfaker

import "github.com/shun-ideguchi/jp-gofaker/internal/dataset"

type CompanyGenerator struct {
	g *Generator
}

func (c CompanyGenerator) Name() string {
	legalTypes := []string{"株式会社", "合同会社"}
	legalType := legalTypes[c.g.rng.Intn(len(legalTypes))]
	prefix := dataset.CompanyPrefixes[c.g.rng.Intn(len(dataset.CompanyPrefixes))]
	core := dataset.CompanyCores[c.g.rng.Intn(len(dataset.CompanyCores))]

	return legalType + prefix + core
}

func (c CompanyGenerator) Department() string {
	values := []string{
		"営業部",
		"開発部",
		"人事部",
		"経理部",
		"カスタマーサポート部",
	}

	return values[c.g.rng.Intn(len(values))]
}

func (c CompanyGenerator) Title() string {
	values := []string{
		"部長",
		"課長",
		"主任",
		"担当者",
		"マネージャー",
	}

	return values[c.g.rng.Intn(len(values))]
}
