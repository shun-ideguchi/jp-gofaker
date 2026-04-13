package jpfaker

import (
	"fmt"

	"github.com/shun-ideguchi/jp-gofaker/internal/dataset"
)

type AddressGenerator struct {
	g *Generator
}

func (a AddressGenerator) PostalCode() string {
	return a.Value().PostalCode
}

func (a AddressGenerator) Prefecture() string {
	return a.Value().Prefecture
}

func (a AddressGenerator) City() string {
	return a.Value().City
}

func (a AddressGenerator) Street() string {
	return a.Value().Street
}

func (a AddressGenerator) Building() string {
	return a.Value().Building
}

func (a AddressGenerator) Full() string {
	return a.Value().Full()
}

func (a AddressGenerator) Value() AddressValue {
	postal := dataset.PostalCodes[a.g.rng.Intn(len(dataset.PostalCodes))]
	streetNames := dataset.StreetNamesByPrefecture[postal.Prefecture]
	streetName := streetNames[a.g.rng.Intn(len(streetNames))]
	buildingName := dataset.BuildingNames[a.g.rng.Intn(len(dataset.BuildingNames))]

	chome := a.g.rng.Intn(5) + 1
	block := a.g.rng.Intn(20) + 1
	number := a.g.rng.Intn(20) + 1
	room := a.g.rng.Intn(8) + 1

	return AddressValue{
		PostalCode: postal.Code,
		Prefecture: postal.Prefecture,
		City:       postal.City,
		Street:     fmt.Sprintf("%s%d-%d-%d", streetName, chome, block, number),
		Building:   fmt.Sprintf("%s%d0%d", buildingName, room, a.g.rng.Intn(9)+1),
	}
}
