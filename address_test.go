package jpfaker_test

import (
	"regexp"
	"strings"
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
	"github.com/shun-ideguchi/jp-gofaker/internal/dataset"
)

func TestAddressValueReturnsPopulatedFields(t *testing.T) {
	// Address.Value が住所を構成する各フィールドを空文字で返さないことを確認する。
	g := jpfaker.New(jpfaker.WithSeed(3))
	address := g.Address().Value()

	if address.PostalCode == "" || address.Prefecture == "" || address.City == "" || address.Street == "" || address.Building == "" {
		t.Fatalf("address fields must not be empty: %+v", address)
	}
}

func TestPostalCodeFormat(t *testing.T) {
	// 郵便番号が NNN-NNNN 形式で生成されることを確認する。
	g := jpfaker.New(jpfaker.WithSeed(4))
	postalCode := g.Address().PostalCode()

	if ok, err := regexp.MatchString(`^\d{3}-\d{4}$`, postalCode); err != nil {
		t.Fatalf("failed to validate postal code: %v", err)
	} else if !ok {
		t.Fatalf("invalid postal code format: %q", postalCode)
	}
}

func TestAddressFullMatchesValue(t *testing.T) {
	// Address.Full が Address.Value().Full() と同じ内容を返すことを確認する。
	valueGenerator := jpfaker.New(jpfaker.WithSeed(5))
	fullGenerator := jpfaker.New(jpfaker.WithSeed(5))
	value := valueGenerator.Address().Value()

	if got, want := fullGenerator.Address().Full(), value.Full(); got != want {
		t.Fatalf("full address mismatch: got %q want %q", got, want)
	}
}

func TestAddressValueUsesMatchingPostalEntry(t *testing.T) {
	// 郵便番号、都道府県、市区町村が同じ郵便データの組み合わせで返ることを確認する。
	address := jpfaker.New(jpfaker.WithSeed(10)).Address().Value()

	matched := false
	for _, postal := range dataset.PostalCodes {
		if postal.Code == address.PostalCode && postal.Prefecture == address.Prefecture && postal.City == address.City {
			matched = true
			break
		}
	}

	if !matched {
		t.Fatalf("address must match a postal dataset entry: %+v", address)
	}
}

func TestStreetUsesPrefectureSpecificPool(t *testing.T) {
	// Street が都道府県ごとの町名候補から生成されることを確認する。
	g := jpfaker.New(jpfaker.WithSeed(9))
	address := g.Address().Value()
	streetNames := dataset.StreetNamesByPrefecture[address.Prefecture]

	if len(streetNames) == 0 {
		t.Fatalf("street pool must exist for prefecture: %q", address.Prefecture)
	}

	if !hasStreetPrefix(address.Street, streetNames) {
		t.Fatalf("street %q must use prefecture specific pool for %q", address.Street, address.Prefecture)
	}
}

func TestAddressBuildingUsesKnownPrefix(t *testing.T) {
	// Building が既知の建物名候補を接頭辞として含むことを確認する。
	address := jpfaker.New(jpfaker.WithSeed(11)).Address().Value()

	found := false
	for _, building := range dataset.BuildingNames {
		if strings.HasPrefix(address.Building, building) {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("building %q must use a known building prefix", address.Building)
	}
}

func hasStreetPrefix(street string, candidates []string) bool {
	for _, candidate := range candidates {
		if strings.HasPrefix(street, candidate) {
			return true
		}
	}

	return false
}
