package jpfaker_test

import (
	"fmt"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func Example() {
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
	// 鈴木 悠真
	// スズキ ユウマ
	// 810-0001
	// 福岡県福岡市中央区天神4-6-1サクラハイツ107
	// 090-8511-8162
	// 合同会社未来物流
	// 開発部
	// 部長
}
