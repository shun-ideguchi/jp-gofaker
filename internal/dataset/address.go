package dataset

type PostalAddress struct {
	Code       string
	Prefecture string
	City       string
}

var PostalCodes = []PostalAddress{
	{Code: "100-0001", Prefecture: "東京都", City: "千代田区"},
	{Code: "060-0001", Prefecture: "北海道", City: "札幌市中央区"},
	{Code: "980-0001", Prefecture: "宮城県", City: "仙台市青葉区"},
	{Code: "231-0001", Prefecture: "神奈川県", City: "横浜市中区"},
	{Code: "460-0001", Prefecture: "愛知県", City: "名古屋市中区"},
	{Code: "530-0001", Prefecture: "大阪府", City: "大阪市北区"},
	{Code: "600-0001", Prefecture: "京都府", City: "京都市下京区"},
	{Code: "810-0001", Prefecture: "福岡県", City: "福岡市中央区"},
}

var StreetNames = []string{
	"桜丘",
	"本町",
	"栄",
	"天神",
	"御幸町",
	"錦",
	"駅前",
	"中央",
}

var BuildingNames = []string{
	"青葉ビル",
	"サクラハイツ",
	"中央レジデンス",
	"みらいタワー",
	"東雲コート",
	"JPスクエア",
}
