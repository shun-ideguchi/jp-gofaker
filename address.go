package jpfaker

import (
	"fmt"

	"github.com/shun-ideguchi/jp-gofaker/internal/dataset"
)

// AddressGenerator generates Japanese postal addresses.
type AddressGenerator struct {
	g *Generator
}

// PostalCode returns a generated postal code.
func (a AddressGenerator) PostalCode() string {
	return a.Value().PostalCode
}

// Prefecture returns a generated prefecture name.
func (a AddressGenerator) Prefecture() string {
	return a.Value().Prefecture
}

// City returns a generated city or ward name.
func (a AddressGenerator) City() string {
	return a.Value().City
}

// Street returns a generated street-level address.
func (a AddressGenerator) Street() string {
	return a.Value().Street
}

// Building returns a generated building and room representation.
func (a AddressGenerator) Building() string {
	return a.Value().Building
}

// Full returns a generated full address string.
func (a AddressGenerator) Full() string {
	return a.Value().Full()
}

// Value returns a generated address broken into structured fields.
func (a AddressGenerator) Value() AddressValue {
	postal, ok := pickOne(a.g.rng, dataset.PostalCodes)
	if !ok {
		return AddressValue{}
	}

	streetName, _ := pickOne(a.g.rng, dataset.StreetNamesByPrefecture[postal.Prefecture])
	buildingName, _ := pickOne(a.g.rng, dataset.BuildingNames)

	chome := a.g.rng.Intn(5) + 1
	block := a.g.rng.Intn(20) + 1
	number := a.g.rng.Intn(20) + 1
	room := a.g.rng.Intn(8) + 1
	building := ""
	if buildingName != "" {
		building = fmt.Sprintf("%s%d0%d", buildingName, room, a.g.rng.Intn(9)+1)
	}

	return AddressValue{
		PostalCode: postal.Code,
		Prefecture: postal.Prefecture,
		City:       postal.City,
		Street:     fmt.Sprintf("%s%d-%d-%d", streetName, chome, block, number),
		Building:   building,
	}
}
