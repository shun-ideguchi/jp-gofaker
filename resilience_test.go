package jpfaker

import (
	"testing"

	"github.com/shun-ideguchi/jp-gofaker/internal/dataset"
)

func TestGeneratorsHandleEmptyDatasets(t *testing.T) {
	originalLastNames := dataset.PersonLastNames
	originalFirstNames := dataset.PersonFirstNames
	originalFirstNamesMale := dataset.PersonFirstNamesMale
	originalFirstNamesFemale := dataset.PersonFirstNamesFemale
	originalFirstNamesNeutral := dataset.PersonFirstNamesNeutral
	originalPostalCodes := dataset.PostalCodes
	originalStreetNames := dataset.StreetNamesByPrefecture
	originalBuildingNames := dataset.BuildingNames
	originalCompanyPrefixes := dataset.CompanyPrefixes
	originalCompanyCores := dataset.CompanyCores
	originalCompanyLegalTypes := dataset.CompanyLegalTypes
	originalCompanyDepartments := dataset.CompanyDepartments
	originalCompanyTitles := dataset.CompanyTitles

	t.Cleanup(func() {
		dataset.PersonLastNames = originalLastNames
		dataset.PersonFirstNames = originalFirstNames
		dataset.PersonFirstNamesMale = originalFirstNamesMale
		dataset.PersonFirstNamesFemale = originalFirstNamesFemale
		dataset.PersonFirstNamesNeutral = originalFirstNamesNeutral
		dataset.PostalCodes = originalPostalCodes
		dataset.StreetNamesByPrefecture = originalStreetNames
		dataset.BuildingNames = originalBuildingNames
		dataset.CompanyPrefixes = originalCompanyPrefixes
		dataset.CompanyCores = originalCompanyCores
		dataset.CompanyLegalTypes = originalCompanyLegalTypes
		dataset.CompanyDepartments = originalCompanyDepartments
		dataset.CompanyTitles = originalCompanyTitles
	})

	dataset.PersonLastNames = nil
	dataset.PersonFirstNames = nil
	dataset.PersonFirstNamesMale = nil
	dataset.PersonFirstNamesFemale = nil
	dataset.PersonFirstNamesNeutral = nil
	dataset.PostalCodes = nil
	dataset.StreetNamesByPrefecture = map[string][]string{}
	dataset.BuildingNames = nil
	dataset.CompanyPrefixes = nil
	dataset.CompanyCores = nil
	dataset.CompanyLegalTypes = nil
	dataset.CompanyDepartments = nil
	dataset.CompanyTitles = nil

	g := New(WithSeed(1))

	if got := g.Person().Name(); got != (PersonName{}) {
		t.Fatalf("empty datasets should return zero person name: %+v", got)
	}

	if got := g.Person().MaleName(); got != (PersonName{}) {
		t.Fatalf("empty male pool should return zero person name: %+v", got)
	}

	if got := g.Person().FemaleName(); got != (PersonName{}) {
		t.Fatalf("empty female pool should return zero person name: %+v", got)
	}

	if got := g.Person().NeutralName(); got != (PersonName{}) {
		t.Fatalf("empty neutral pool should return zero person name: %+v", got)
	}

	if got := g.Address().Value(); got != (AddressValue{}) {
		t.Fatalf("empty address datasets should return zero address: %+v", got)
	}

	if got := g.Company().Name(); got != "" {
		t.Fatalf("empty company datasets should return empty name: %q", got)
	}

	if got := g.Company().Department(); got != "" {
		t.Fatalf("empty department datasets should return empty string: %q", got)
	}

	if got := g.Company().Title(); got != "" {
		t.Fatalf("empty title datasets should return empty string: %q", got)
	}
}

func TestAddressValueHandlesMissingStreetAndBuildingPools(t *testing.T) {
	originalPostalCodes := dataset.PostalCodes
	originalStreetNames := dataset.StreetNamesByPrefecture
	originalBuildingNames := dataset.BuildingNames

	t.Cleanup(func() {
		dataset.PostalCodes = originalPostalCodes
		dataset.StreetNamesByPrefecture = originalStreetNames
		dataset.BuildingNames = originalBuildingNames
	})

	dataset.PostalCodes = []dataset.PostalAddress{
		{Code: "100-0001", Prefecture: "東京都", City: "千代田区"},
	}
	dataset.StreetNamesByPrefecture = map[string][]string{}
	dataset.BuildingNames = nil

	got := New(WithSeed(1)).Address().Value()

	if got.PostalCode != "100-0001" || got.Prefecture != "東京都" || got.City != "千代田区" {
		t.Fatalf("address should preserve postal information when optional pools are missing: %+v", got)
	}

	if got.Building != "" {
		t.Fatalf("building should be empty when building pool is missing: %q", got.Building)
	}
}

func TestCompanyNameUsesAvailableParts(t *testing.T) {
	originalCompanyPrefixes := dataset.CompanyPrefixes
	originalCompanyCores := dataset.CompanyCores
	originalCompanyLegalTypes := dataset.CompanyLegalTypes

	t.Cleanup(func() {
		dataset.CompanyPrefixes = originalCompanyPrefixes
		dataset.CompanyCores = originalCompanyCores
		dataset.CompanyLegalTypes = originalCompanyLegalTypes
	})

	dataset.CompanyPrefixes = nil
	dataset.CompanyCores = []string{"システム"}
	dataset.CompanyLegalTypes = []string{"株式会社"}

	if got := New(WithSeed(1)).Company().Name(); got != "株式会社システム" {
		t.Fatalf("company name should be built from available parts: %q", got)
	}
}
