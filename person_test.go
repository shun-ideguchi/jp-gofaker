package jpfaker_test

import (
	"strings"
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
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
	fullNameGenerator := jpfaker.New(jpfaker.WithSeed(2))
	fullNameKanaGenerator := jpfaker.New(jpfaker.WithSeed(2))
	separatorGenerator := jpfaker.New(jpfaker.WithSeed(2))
	kanaSeparatorGenerator := jpfaker.New(jpfaker.WithSeed(2))

	name := nameGenerator.Person().Name()

	if got, want := fullNameGenerator.Person().FullName(), name.FullName(); got != want {
		t.Fatalf("full name mismatch: got %q want %q", got, want)
	}

	if got, want := fullNameKanaGenerator.Person().FullNameKana(), name.FullNameKana(); got != want {
		t.Fatalf("full name kana mismatch: got %q want %q", got, want)
	}

	fullName := separatorGenerator.Person().FullName()
	if !strings.Contains(fullName, " ") {
		t.Fatalf("full name must contain a separator: %q", fullName)
	}

	fullNameKana := kanaSeparatorGenerator.Person().FullNameKana()
	if !strings.Contains(fullNameKana, " ") {
		t.Fatalf("full name kana must contain a separator: %q", fullNameKana)
	}
}
