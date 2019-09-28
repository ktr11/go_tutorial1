// スライス[前編]
package main

import (
	"fmt"
)

func main() {
	// スライスのもととなる配列を作成
	num := [5]int{0, 1, 2, 3, 4}
	// スライス型の変数の宣言
	var slice1 []int

	//配列全体
	slice1 = num[:]
	fmt.Println(slice1)
	// インデックス1から4より前まで
	slice2 := num[1:4]
	fmt.Println(slice2)
	// インデックス4から
	slice3 := num[4:]
	fmt.Println(slice3)
	// インデックス4より前まで
	slice4 := num[:4]
	fmt.Println(slice4)

	// スライスのもととなる配列を作成
	num2 := [5]int{0, 1, 2, 3, 4}
	// 配列をスライスとして関数に渡す
	plusOne(num2[:])
	fmt.Println(num2)

	// スライスのもととなる配列を作成
	num3 := [5]int{0, 1, 2, 3, 4}

	// 配列の一部をスライス
	sliceA := num3[1:4]
	fmt.Println("sliceA=", sliceA)
	fmt.Println("len=", len(sliceA))
	fmt.Println("cap=", cap(sliceA))

	// スライスの一部をスライス
	sliceB := sliceA[1:4] // 長さを超過したキャパシティ最大値まで取得可能
	fmt.Println("sliceB=", sliceB)
	fmt.Println("len=", len(sliceB))
	fmt.Println("cap=", cap(sliceB))
}

// 要素内の数字をすべて +1 する
func plusOne(vals []int) {
	for i, m := range vals {
		vals[i] = m + 1
	}
}
