package main

import "fmt"

// LoopNum レシーバ変数 簡単な例
type LoopNum int

// SavingBox オブジェクト指向らしくプログラムを扱うことができる
type SavingBox struct {
	name  string
	money int
}

// NewBox SavingBox作成
func NewBox(name string) *SavingBox {
	var sv = new(SavingBox)
	sv.name = name
	return sv
}

// Income svにお金を入れる。
func (sv *SavingBox) Income(ammount int) {
	sv.money += ammount
}

// Break svを壊す
func (sv *SavingBox) Break() int {
	lastMoney := sv.money
	sv.money = 0
	return lastMoney
}

func main() {
	DisplayHello()
	DisplaySum(17, 99)
	DisplaySumAll(17, 99, 11, 22)
	fmt.Println(SumAll(1, 2, 3))

	syou, amari := Div(19, 4)
	fmt.Printf("%dあまり%dです。\n", syou, amari)

	var lpNum LoopNum = 3
	lpNum.TimesDisplay("Hello")

	box := NewBox("貯金箱A")
	box.Income(100)
	box.Income(200)
	box.Income(500)
	fmt.Printf("%sを壊したら%d円出てきました。", box.name, box.Break())

}

// DisplayHello Helloを表示
func DisplayHello() {
	fmt.Println("Hello!")
}

// DisplaySum 引数に指定したint型の変数を足す
func DisplaySum(left int, right int) {
	fmt.Println(left + right)
}

// DisplaySumAll 型名の前に「...」を用いて可変長引数
// 関数内では配列のように動作する。
func DisplaySumAll(values ...int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println(sum)
}

// SumAll 戻り値を指定
func SumAll(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

// Div 戻り値を複数指定することも可能
func Div(left int, right int) (int, int) {
	return left / right, left % right
}

// TimesDisplay レシーバ変数とメソッド
// 関数名の前にレシーバ変数を指定することができる。
// レシーバ変数を指定した関数のことを、その変数のメソッドと呼ぶ
func (n LoopNum) TimesDisplay(s string) {
	for i := 0; i < int(n); i++ {
		fmt.Println(s)
	}
}
