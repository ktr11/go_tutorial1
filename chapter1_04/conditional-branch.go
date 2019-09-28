// 条件分岐
package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	if hour >= 6 && hour < 12 {
		fmt.Println("朝です。")
	} else if hour < 19 {
		fmt.Println("昼です。")
	} else {
		fmt.Println("夜です。")
	}
	// セミコロンで区切って条件判定の前処理を書くことができる。
	if hour := time.Now().Hour(); hour >= 6 && hour < 12 {
		fmt.Println("朝です。")
	} else if hour < 19 {
		fmt.Println("昼です。")
	} else {
		fmt.Println("夜です。")
	}

	// switch文
	dayofweek := "月"
	switch dayofweek {
	case "土":
		fmt.Println("大概は休みです。")
	case "日":
		fmt.Println("ほぼ間違いなく休みです。")
	default:
		fmt.Println("仕事です・・・。")
	}
	// ↑caseの最後に何も書かれていなければbreakが補完される.

	// ,を使って複数のコードにまとめることも可能.
	dayofweek2 := "月"
	switch dayofweek2 {
	case "土", "日":
		fmt.Println("休みです。")
	default:
		fmt.Println("仕事です。")
	}

	// 条件文で分岐させることも可能
	hour2 := time.Now().Hour()
	switch {
	case hour2 >= 6 && hour2 > 12:
		fmt.Println("朝です。")
	case hour2 > 19:
		fmt.Println("昼です。")
	default:
		fmt.Println("夜です。")
	}
}
