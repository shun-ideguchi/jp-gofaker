package jpfaker_test

import (
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func TestWithSeedProducesDeterministicValues(t *testing.T) {
	// 同じ seed から生成した各ドメインの値が再現されることを確認する。
	left := jpfaker.New(jpfaker.WithSeed(42))
	right := jpfaker.New(jpfaker.WithSeed(42))

	tests := []struct {
		name string
		got  func() any
		want func() any
	}{
		{name: "person", got: func() any { return left.Person().Name() }, want: func() any { return right.Person().Name() }},
		{name: "address", got: func() any { return left.Address().Value() }, want: func() any { return right.Address().Value() }},
		{name: "mobile", got: func() any { return left.Phone().Mobile() }, want: func() any { return right.Phone().Mobile() }},
		{name: "landline", got: func() any { return left.Phone().Landline() }, want: func() any { return right.Phone().Landline() }},
		{name: "toll free", got: func() any { return left.Phone().TollFree() }, want: func() any { return right.Phone().TollFree() }},
		{name: "company name", got: func() any { return left.Company().Name() }, want: func() any { return right.Company().Name() }},
		{name: "department", got: func() any { return left.Company().Department() }, want: func() any { return right.Company().Department() }},
		{name: "title", got: func() any { return left.Company().Title() }, want: func() any { return right.Company().Title() }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := tt.got(), tt.want(); got != want {
				t.Fatalf("%s mismatch: got %#v want %#v", tt.name, got, want)
			}
		})
	}
}
