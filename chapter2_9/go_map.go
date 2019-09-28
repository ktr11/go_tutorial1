// マップ
package main

import (
	"fmt"
)

func main() {
	// マップの作成
	currencies := make(map[string]string)
	// キーを指定して値を格納
	currencies["日本"] = "JPY"
	currencies["USA"] = "USD"
	currencies["EU"] = "EUR"
	currencies["中国"] = "CHY"
	// キーを指定して値を取得
	fmt.Println(currencies["日本"])
	fmt.Println(currencies["中国"])

	fmt.Println("== すべて取得 =======")
	// for,rangeを使用してマップ内の値をすべて取得
	for country, currency := range currencies {
		fmt.Println(country, currency)
	}
	fmt.Println("== ここまで　 =======")

	// キーの存在確認
	currency, exist := currencies["ロシア"]
	fmt.Println(currency)
	fmt.Println(exist)
	if exist {
		fmt.Println("登録済み", currency)
	} else {
		fmt.Println("未登録")
	}

	// キーの削除
	// マップから削除するには「delete」関数を使用します
	fmt.Println("「delete」関数を使用")
	delete(currencies, "日本")
	for country, currency := range currencies {
		fmt.Println(country, currency)
	}

	// マップの初期化方法
	fmt.Println("マップの初期化")
	currencies2 := map[string]string{
		"日本": "JPY",
		"米国": "USD",
		"EU": "EUR",
		"中国": "CNY",
	}
	for country, currency := range currencies2 {
		fmt.Println(country, currency)
	}

}
