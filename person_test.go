package jpfaker_test

import (
	"slices"
	"strings"
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
	"github.com/shun-ideguchi/jp-gofaker/internal/dataset"
)

func TestPersonNameReturnsPopulatedFields(t *testing.T) {
	// Person.Name が姓名とかなの各フィールドを空文字で返さないことを確認する。
	g := jpfaker.New(jpfaker.WithSeed(1))
	name := g.Person().Name()

	if name.LastName == "" || name.FirstName == "" {
		t.Fatalf("name must not be empty: %+v", name)
	}

	if name.LastNameKana == "" || name.FirstNameKana == "" {
		t.Fatalf("kana must not be empty: %+v", name)
	}
}

func TestPersonFullNameHelpers(t *testing.T) {
	// FullName 系ヘルパーが Name の内容と一致し、区切りを含むことを確認する。
	nameGenerator := jpfaker.New(jpfaker.WithSeed(2))
	name := nameGenerator.Person().Name()
	tests := []struct {
		name string
		got  func() string
		want string
	}{
		{name: "full name", got: func() string { return jpfaker.New(jpfaker.WithSeed(2)).Person().FullName() }, want: name.FullName()},
		{name: "full name kana", got: func() string { return jpfaker.New(jpfaker.WithSeed(2)).Person().FullNameKana() }, want: name.FullNameKana()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.got()
			if got != tt.want {
				t.Fatalf("%s mismatch: got %q want %q", tt.name, got, tt.want)
			}
			if !strings.Contains(got, " ") {
				t.Fatalf("%s must contain a separator: %q", tt.name, got)
			}
		})
	}
}

func TestPersonNameVariantsUseExpectedPools(t *testing.T) {
	// MaleName、FemaleName、NeutralName が対応する名データ集合から生成することを確認する。
	g := jpfaker.New(jpfaker.WithSeed(8))

	male := g.Person().MaleName()
	female := g.Person().FemaleName()
	neutral := g.Person().NeutralName()

	if !containsName(dataset.PersonFirstNamesMale, male.FirstName, male.FirstNameKana) {
		t.Fatalf("male name must be chosen from male pool: %+v", male)
	}

	if !containsName(dataset.PersonFirstNamesFemale, female.FirstName, female.FirstNameKana) {
		t.Fatalf("female name must be chosen from female pool: %+v", female)
	}

	if !containsName(dataset.PersonFirstNamesNeutral, neutral.FirstName, neutral.FirstNameKana) {
		t.Fatalf("neutral name must be chosen from neutral pool: %+v", neutral)
	}
}

func containsName(pool []dataset.Name, firstName, firstNameKana string) bool {
	return slices.ContainsFunc(pool, func(n dataset.Name) bool {
		return n.Text == firstName && n.Kana == firstNameKana
	})
}
