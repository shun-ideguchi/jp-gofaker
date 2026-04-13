package jpfaker_test

import (
	"fmt"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func Example() {
	// 公開 API の基本的な利用例と出力例を示す。
	g := jpfaker.New(jpfaker.WithSeed(1))

	name := g.Person().Name()
	address := g.Address().Value()

	fmt.Println(name.FullName())
	fmt.Println(name.FullNameKana())
	fmt.Println(address.PostalCode)
	fmt.Println(address.Full())
	fmt.Println(g.Phone().Mobile())
	fmt.Println(g.Company().Name())
	fmt.Println(g.Company().Department())
	fmt.Println(g.Company().Title())

	// Output:
	// 鈴木 咲良
	// スズキ サクラ
	// 500-8844
	// 岐阜県岐阜市吉野町橋本町4-6-1サクラハイツ107
	// 090-8511-8162
	// 有限会社西日本製作所
	// 経理部
	// 係長
}
