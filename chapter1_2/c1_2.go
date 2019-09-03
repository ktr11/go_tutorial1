// 型の宣言と変換について
package main

import "fmt"

// Score typeキーワードで型を宣言
type Score int

func main() {

	var myScore Score = 100
	fmt.Printf("私の点数は%d点です。\n", myScore)
}
