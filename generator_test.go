package jpfaker_test

import (
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func TestWithSeedProducesDeterministicValues(t *testing.T) {
	// 同じ seed から生成した各ドメインの値が再現されることを確認する。
	left := jpfaker.New(jpfaker.WithSeed(42))
	right := jpfaker.New(jpfaker.WithSeed(42))

	if got, want := left.Person().Name(), right.Person().Name(); got != want {
		t.Fatalf("person mismatch: got %+v want %+v", got, want)
	}

	if got, want := left.Address().Value(), right.Address().Value(); got != want {
		t.Fatalf("address mismatch: got %+v want %+v", got, want)
	}

	if got, want := left.Phone().Mobile(), right.Phone().Mobile(); got != want {
		t.Fatalf("mobile mismatch: got %q want %q", got, want)
	}

	if got, want := left.Phone().Landline(), right.Phone().Landline(); got != want {
		t.Fatalf("landline mismatch: got %q want %q", got, want)
	}

	if got, want := left.Phone().TollFree(), right.Phone().TollFree(); got != want {
		t.Fatalf("toll free mismatch: got %q want %q", got, want)
	}

	if got, want := left.Company().Name(), right.Company().Name(); got != want {
		t.Fatalf("company name mismatch: got %q want %q", got, want)
	}

	if got, want := left.Company().Department(), right.Company().Department(); got != want {
		t.Fatalf("department mismatch: got %q want %q", got, want)
	}

	if got, want := left.Company().Title(), right.Company().Title(); got != want {
		t.Fatalf("title mismatch: got %q want %q", got, want)
	}
}
