package jpfaker_test

import (
	"regexp"
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
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
