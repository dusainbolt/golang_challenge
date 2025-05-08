package main

import (
	"encoding/json"
	"fmt"
)

// Stock đại diện cho một cổ phiếu với mã và tỷ lệ cổ tức
type Stock struct {
	Ticker   string  `json:"ticker"`   // mã cổ phiếu
	Dividend float64 `json:"dividend"` // tỷ lệ cổ tức
}

// HighestDividend nhận vào một chuỗi JSON chứa danh sách cổ phiếu,
// giải mã (unmarshal) thành mảng Stock,
// sau đó tìm và trả về mã cổ phiếu có cổ tức cao nhất.
func HighestDividend(str string) string {
	var ticker []Stock

	// Chuyển chuỗi JSON sang slice các struct Stock
	if err := json.Unmarshal([]byte(str), &ticker); err != nil {
		panic(err) // nếu lỗi khi parse JSON thì dừng chương trình
	}

	// Nếu danh sách rỗng thì trả về chuỗi rỗng
	if len(ticker) == 0 {
		return ""
	}

	// Gán cổ phiếu đầu tiên làm "cao nhất" tạm thời
	highest := ticker[0]

	// Duyệt qua từng cổ phiếu để tìm cổ tức cao nhất
	for _, v := range ticker {
		if highest.Dividend < v.Dividend {
			highest = v // cập nhật cổ phiếu nếu cổ tức cao hơn
		}
	}

	// Trả về mã của cổ phiếu có cổ tức cao nhất
	return highest.Ticker
}

func main() {
	fmt.Println("Stock Price AI")

	stocks_json := `[
		{"ticker":"APPL","dividend":0.5},
		{"ticker":"GOOG","dividend":0.2},
		{"ticker":"DUMAX","dividend":11.2},
		{"ticker":"MSFT","dividend":0.8},
		{"ticker":"MSFX","dividend":3}
	]`

	highestDividend := HighestDividend(stocks_json)
	fmt.Println(highestDividend)
}
