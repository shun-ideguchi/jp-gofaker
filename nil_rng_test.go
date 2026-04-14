package jpfaker

import (
	"regexp"
	"testing"
)

func TestNewFallsBackWhenRandomSourceIsNil(t *testing.T) {
	g := New(withRandomSource(nil))

	if got := g.Person().Name(); got == (PersonName{}) {
		t.Fatal("person name should still be generated when rng is nil")
	}

	if got := g.Address().Value(); got == (AddressValue{}) {
		t.Fatal("address should still be generated when rng is nil")
	}

	if got := g.Company().Name(); got == "" {
		t.Fatal("company name should still be generated when rng is nil")
	}

	assertMatches := func(pattern, value string) {
		t.Helper()

		ok, err := regexp.MatchString(pattern, value)
		if err != nil {
			t.Fatalf("failed to validate %q with %q: %v", value, pattern, err)
		}
		if !ok {
			t.Fatalf("value %q does not match %q", value, pattern)
		}
	}

	assertMatches(`^0[789]0-\d{4}-\d{4}$`, g.Phone().Mobile())
	assertMatches(`^0\d{1,3}-\d{3,4}-\d{4}$`, g.Phone().Landline())
	assertMatches(`^(0120|0800)-\d{3}-\d{3}$`, g.Phone().TollFree())
}
