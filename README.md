# jp-gofaker

`jp-gofaker` は Go アプリケーション向けの日本語ダミーデータ生成パッケージです。

テスト、seed データ、ローカル検証、デモ用途を想定しています。`WithSeed` を使うと再現可能な値を生成できます。

## インストール

```bash
go get github.com/shun-ideguchi/jp-gofaker
```

## 使い方

```go
package main

import (
	"fmt"

	jpfaker "github.com/shun-ideguchi/jp-gofaker"
)

func main() {
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
}
```

出力例:

```text
鈴木 咲良
スズキ サクラ
500-8844
岐阜県岐阜市吉野町橋本町4-6-1サクラハイツ107
090-8511-8162
有限会社西日本製作所
経理部
係長
```

## 生成できる値

### Person

- `LastName()`
- `FirstName()`
- `FullName()`
- `LastNameKana()`
- `FirstNameKana()`
- `FullNameKana()`
- `Name()`
- `MaleName()`
- `FemaleName()`
- `NeutralName()`

### Address

- `PostalCode()`
- `Prefecture()`
- `City()`
- `Street()`
- `Building()`
- `Full()`
- `Value()`

### Phone

- `Mobile()`
- `Landline()`
- `TollFree()`

### Company

- `Name()`
- `Department()`
- `Title()`

## `WithSeed` について

`WithSeed(seed)` を渡すと、同じ seed から同じ順序で同じ値が生成されます。テストやスナップショット比較に向いています。

```go
left := jpfaker.New(jpfaker.WithSeed(42))
right := jpfaker.New(jpfaker.WithSeed(42))

fmt.Println(left.Person().FullName() == right.Person().FullName())
// true
```

## 制約

- 値は実在人物や実在企業を表すものではなく、あくまでダミーデータです
- 住所や企業名は自然な見た目を優先した合成値であり、実在性を保証しません
- 現時点では locale 切り替えや出力フォーマット切り替えには対応していません

## 開発

```bash
go test ./...
```
