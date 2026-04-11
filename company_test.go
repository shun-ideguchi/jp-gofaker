package jpfaker_test

import (
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func TestCompanyValuesAreNotEmpty(t *testing.T) {
	// 会社名、部署名、役職名が空文字で生成されないことを確認する。
	g := jpfaker.New(jpfaker.WithSeed(7))

	if got := g.Company().Name(); got == "" {
		t.Fatal("company name must not be empty")
	}

	if got := g.Company().Department(); got == "" {
		t.Fatal("department must not be empty")
	}

	if got := g.Company().Title(); got == "" {
		t.Fatal("title must not be empty")
	}
}
