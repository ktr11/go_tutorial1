// 繰り返し
package main

import "fmt"

func main() {
	// 一般的なfor
	for i := 1; i < 100; i++ {
		if i/2 != 0 {
			fmt.Println(i)
		}
	}

	// 無限ループ
	j := 0
	for {
		j++
		if j >= 100 {
			break
		} else if j/2 == 0 {
			continue
		}
		fmt.Println(j)
	}
	// コレクション内の要素のイテレーション
	// rangeキーワードを使用することで、配列をはじめとするコレクション内の要素を取り出して処理をすることが可能
	dayOfWeeks := [...]string{"月", "火", "水", "木", "金", "土", "日"}
	for arrayIndex, dayOfWeek := range dayOfWeeks {
		fmt.Printf("%d番目の曜日は%s曜日です。\n", arrayIndex+1, dayOfWeek)
	}
}
