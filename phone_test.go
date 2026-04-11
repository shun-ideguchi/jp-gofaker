package jpfaker_test

import (
	"regexp"
	"testing"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func TestPhoneFormats(t *testing.T) {
	// 携帯電話、固定電話、フリーダイヤルがそれぞれ想定形式で生成されることを確認する。
	g := jpfaker.New(jpfaker.WithSeed(6))

	assertMatches(t, `^0[789]0-\d{4}-\d{4}$`, g.Phone().Mobile())
	assertMatches(t, `^0\d{1,3}-\d{3,4}-\d{4}$`, g.Phone().Landline())
	assertMatches(t, `^(0120|0800)-\d{3}-\d{3}$`, g.Phone().TollFree())
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
