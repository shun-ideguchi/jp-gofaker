package jpfaker_test

import (
	"regexp"
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func TestPhoneFormats(t *testing.T) {
	// 携帯電話、固定電話、フリーダイヤルがそれぞれ想定形式で生成されることを確認する。
	g := jpfaker.New(jpfaker.WithSeed(6))
	tests := []struct {
		name    string
		pattern string
		value   string
	}{
		{name: "mobile", pattern: `^0[789]0-\d{4}-\d{4}$`, value: g.Phone().Mobile()},
		{name: "landline", pattern: `^0\d{1,3}-\d{3,4}-\d{4}$`, value: g.Phone().Landline()},
		{name: "toll free", pattern: `^(0120|0800)-\d{3}-\d{3}$`, value: g.Phone().TollFree()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertMatches(t, tt.pattern, tt.value)
		})
	}
}

func assertMatches(t *testing.T, pattern, value string) {
	// 与えた文字列が期待する正規表現に一致することを確認する。
	t.Helper()

	ok, err := regexp.MatchString(pattern, value)
	if err != nil {
		t.Fatalf("failed to validate %q with %q: %v", value, pattern, err)
	}
	if !ok {
		t.Fatalf("value %q does not match %q", value, pattern)
	}
}
