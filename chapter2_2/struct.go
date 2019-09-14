// 構造体
package main

import (
	"fmt"

	"github.com/ktr11/go_tutorial1/chapter2_2/test"
)

type Vector struct {
	X int
	Y int
}

func main() {
	var v Vector

	v.X = 2
	v.Y = 5
	fmt.Println(v)
	// 1度に初期化も可能
	w := Vector{X: 3, Y: 4}
	fmt.Println(w)

	var v2 test.Vector3
	v2.X = 2
	v2.Y = 5
	// v2.z = 8 // コンパイルエラー
	// Javaのsetterみたいな関数を作ったので、それを使用
	test.SetZ(&v2, 8)
	fmt.Println(v2)

	fmt.Println(test.NewVector())

	var v3 = test.NewVector()
	v3.X = 1
	v3.Y = 2
	v3.Z = 3
	fmt.Println(v3)
}
