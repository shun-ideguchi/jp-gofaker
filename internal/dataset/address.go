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
	{Code: "330-0854", Prefecture: "埼玉県", City: "さいたま市大宮区"},
	{Code: "261-0023", Prefecture: "千葉県", City: "千葉市美浜区"},
	{Code: "390-0811", Prefecture: "長野県", City: "松本市中央"},
	{Code: "420-0858", Prefecture: "静岡県", City: "静岡市葵区"},
	{Code: "460-0001", Prefecture: "愛知県", City: "名古屋市中区"},
	{Code: "500-8844", Prefecture: "岐阜県", City: "岐阜市吉野町"},
	{Code: "530-0001", Prefecture: "大阪府", City: "大阪市北区"},
	{Code: "600-0001", Prefecture: "京都府", City: "京都市下京区"},
	{Code: "650-0021", Prefecture: "兵庫県", City: "神戸市中央区"},
	{Code: "700-0901", Prefecture: "岡山県", City: "岡山市北区"},
	{Code: "730-0011", Prefecture: "広島県", City: "広島市中区"},
	{Code: "760-0023", Prefecture: "香川県", City: "高松市寿町"},
	{Code: "790-0003", Prefecture: "愛媛県", City: "松山市三番町"},
	{Code: "810-0001", Prefecture: "福岡県", City: "福岡市中央区"},
	{Code: "850-0058", Prefecture: "長崎県", City: "長崎市尾上町"},
	{Code: "860-0805", Prefecture: "熊本県", City: "熊本市中央区"},
	{Code: "890-0053", Prefecture: "鹿児島県", City: "鹿児島市中央町"},
	{Code: "900-0015", Prefecture: "沖縄県", City: "那覇市久茂地"},
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
	"桜木町",
	"青葉",
	"若葉",
	"緑町",
	"南町",
	"北町",
	"東町",
	"西町",
	"旭町",
	"泉町",
	"高砂",
	"大手町",
	"丸の内",
	"今泉",
	"春日",
	"三条",
	"花園",
	"富士見",
	"豊洲",
	"みなとみらい",
}

var BuildingNames = []string{
	"青葉ビル",
	"サクラハイツ",
	"中央レジデンス",
	"みらいタワー",
	"東雲コート",
	"JPスクエア",
	"グリーンハイツ",
	"サンライズビル",
	"パークサイドレジデンス",
	"ヒルズコート",
	"アーバンスクエア",
	"リバーサイドタワー",
	"メゾン中央",
	"エクセル本町",
	"フォレストテラス",
	"ベイフロントタワー",
}
