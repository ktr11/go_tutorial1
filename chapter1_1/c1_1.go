package main

import "fmt"

// 基本的な文法
func main() {

	// 変数宣言は型名を変数名の後ろに書く。
	var num int
	var pow int
	// 宣言と初期化を同時に行う場合は型の記述は不要
	var result = 1
	// 1行で文が終わる場合セミコロン不要
	num = 2
	pow = 4
	// if,forなどのカッコは不要
	// 1行に文が複数の場合はセミコロンが必要
	for i := 0; i < pow; i++ {
		result *= num
	} // ブロック内が1文でも波カッコは必要

	fmt.Printf("%dの%d乗は%dです。\n", num, pow, result)
}
