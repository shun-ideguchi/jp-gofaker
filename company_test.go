package jpfaker_test

import (
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func TestCompanyValuesAreNotEmpty(t *testing.T) {
	// 会社名、部署名、役職名が空文字で生成されないことを確認する。
	g := jpfaker.New(jpfaker.WithSeed(7))
	tests := []struct {
		name  string
		value string
	}{
		{name: "company name", value: g.Company().Name()},
		{name: "department", value: g.Company().Department()},
		{name: "title", value: g.Company().Title()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value == "" {
				t.Fatalf("%s must not be empty", tt.name)
			}
		})
	}
}
