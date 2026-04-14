package jpfaker

import "github.com/shun-ideguchi/jp-gofaker/internal/dataset"

// PersonGenerator generates Japanese personal names.
type PersonGenerator struct {
	g *Generator
}

// LastName returns a generated surname.
func (p PersonGenerator) LastName() string {
	return p.Name().LastName
}

// FirstName returns a generated given name.
func (p PersonGenerator) FirstName() string {
	return p.Name().FirstName
}

// FullName returns a generated full name.
func (p PersonGenerator) FullName() string {
	return p.Name().FullName()
}

// LastNameKana returns a generated surname in kana.
func (p PersonGenerator) LastNameKana() string {
	return p.Name().LastNameKana
}

// FirstNameKana returns a generated given name in kana.
func (p PersonGenerator) FirstNameKana() string {
	return p.Name().FirstNameKana
}

// FullNameKana returns a generated full name in kana.
func (p PersonGenerator) FullNameKana() string {
	return p.Name().FullNameKana()
}

// Name returns a generated personal name with kana.
func (p PersonGenerator) Name() PersonName {
	last, ok := pickOne(p.g.rng, dataset.PersonLastNames)
	if !ok {
		return PersonName{}
	}

	first, ok := pickOne(p.g.rng, dataset.PersonFirstNames)
	if !ok {
		return PersonName{}
	}

	return PersonName{
		LastName:      last.Text,
		FirstName:     first.Text,
		LastNameKana:  last.Kana,
		FirstNameKana: first.Kana,
	}
}

// MaleName returns a generated name using the male first-name pool.
func (p PersonGenerator) MaleName() PersonName {
	return p.nameFrom(dataset.PersonFirstNamesMale)
}

// FemaleName returns a generated name using the female first-name pool.
func (p PersonGenerator) FemaleName() PersonName {
	return p.nameFrom(dataset.PersonFirstNamesFemale)
}

// NeutralName returns a generated name using the neutral first-name pool.
func (p PersonGenerator) NeutralName() PersonName {
	return p.nameFrom(dataset.PersonFirstNamesNeutral)
}

func (p PersonGenerator) nameFrom(firstNames []dataset.Name) PersonName {
	last, ok := pickOne(p.g.rng, dataset.PersonLastNames)
	if !ok {
		return PersonName{}
	}

	first, ok := pickOne(p.g.rng, firstNames)
	if !ok {
		return PersonName{}
	}

	return PersonName{
		LastName:      last.Text,
		FirstName:     first.Text,
		LastNameKana:  last.Kana,
		FirstNameKana: first.Kana,
	}
}
