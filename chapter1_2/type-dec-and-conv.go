// 型の宣言と変換について
package main

import "fmt"

// Score typeキーワードで型を宣言
type Score int

// show
// 関数のレシーバ型としての使用
func (s Score) show() { fmt.Printf("私の点数は%d点です。\n", s) }

// Dictionary 宣言
type Dictionary struct {
	name    string
	meaning string
}

// ReadFunc 宣言
type ReadFunc func(Dictionary) string

func main() {
	var myScore Score = 100
	myScore.show()

	var readFunc ReadFunc
	var dict Dictionary
	readFunc = readOut
	dict.name = "コーヒー"
	dict.meaning = "コーヒー豆から作られる黒色の飲み物"
	fmt.Println(readFunc(dict))

	// 型変換を行うときは、変数を丸カッコで囲み、手前に型名を書く
	var myScore2 Score = 100
	showInt(int(myScore2))
}

func showInt(i int) {
	fmt.Printf("value: %d\n", i)
}

func readOut(d Dictionary) string {
	return fmt.Sprintf("「%s」 は 「%s」 という意味です", d.name, d.meaning)
}
