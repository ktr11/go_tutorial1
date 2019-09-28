// 配列
package main

import "fmt"

func main() {
	// 配列の宣言方法
	// [長さ]型
	var a [1]byte                 // 1byteのbyte型配列
	var b [10]int                 // 10byteのint型配列
	var c [12]struct{ e, f rune } // 構造体の配列
	var d [2][8]rune              // 二次元配列
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(len(c))
	fmt.Println(len(d))

	// 配列の宣言
	var month [12]string
	month[0] = "January"
	month[1] = "Fubruary"
	month[2] = "March"
	month[3] = "April"
	month[4] = "May"
	month[5] = "June"
	month[6] = "July"
	month[7] = "August"
	month[8] = "September"
	month[9] = "October"
	month[10] = "Nobember"
	month[11] = "December"

	// 配列の長さの回数、ループして値を表示します。
	for i := 0; i < len(month); i++ {
		fmt.Printf("%d月 = %s\n", i+1, month[i])
	}
	// 配列にアクセスするfor文では、rangeを使用してアクセスすることも可能です。
	for i, m := range month {
		fmt.Printf("%d月 = %s\n", i+1, m)
	}
	// 配列に直接アクセスする場合、長さを超えたインデックスにアクセスするとランタイムパニックが発生する。
	/*
		var panicArray [12]string
		i := 12
		panicArray[i] = "Undecimber"
		→panic: runtime error: index out of range
	*/
	// 配列の宣言とデータ格納とを同時に行うにはいろいろな書き方がある。
	// 長さ5のint型配列。各要素すべて0
	arr1 := [5]int64{}
	// arr1と同じだが、指定していない要素はすべて0
	arr2 := [5]int64{1, 2, 3}
	// 長さに...を指定すると、指定した要素数が長さとして使用される。
	arr3 := [...]string{"One", "Two", "Three"}
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	fmt.Println(len(arr3))
}
