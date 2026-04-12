package jpfaker

import "github.com/shun-ideguchi/jp-gofaker/internal/dataset"

type PersonGenerator struct {
	g *Generator
}

func (p PersonGenerator) LastName() string {
	return p.Name().LastName
}

func (p PersonGenerator) FirstName() string {
	return p.Name().FirstName
}

func (p PersonGenerator) FullName() string {
	return p.Name().FullName()
}

func (p PersonGenerator) LastNameKana() string {
	return p.Name().LastNameKana
}

func (p PersonGenerator) FirstNameKana() string {
	return p.Name().FirstNameKana
}

func (p PersonGenerator) FullNameKana() string {
	return p.Name().FullNameKana()
}

func (p PersonGenerator) Name() PersonName {
	last := dataset.PersonLastNames[p.g.rng.Intn(len(dataset.PersonLastNames))]
	first := dataset.PersonFirstNames[p.g.rng.Intn(len(dataset.PersonFirstNames))]

	return PersonName{
		LastName:      last.Text,
		FirstName:     first.Text,
		LastNameKana:  last.Kana,
		FirstNameKana: first.Kana,
	}
}

func (p PersonGenerator) MaleName() PersonName {
	return p.nameFrom(dataset.PersonFirstNamesMale)
}

func (p PersonGenerator) FemaleName() PersonName {
	return p.nameFrom(dataset.PersonFirstNamesFemale)
}

func (p PersonGenerator) NeutralName() PersonName {
	return p.nameFrom(dataset.PersonFirstNamesNeutral)
}

func (p PersonGenerator) nameFrom(firstNames []dataset.Name) PersonName {
	last := dataset.PersonLastNames[p.g.rng.Intn(len(dataset.PersonLastNames))]
	first := firstNames[p.g.rng.Intn(len(firstNames))]

	return PersonName{
		LastName:      last.Text,
		FirstName:     first.Text,
		LastNameKana:  last.Kana,
		FirstNameKana: first.Kana,
	}
}
