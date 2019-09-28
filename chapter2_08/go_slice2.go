// スライス[後編]
package main

import (
	"fmt"
)

func main() {
	// スライスを作成
	s1 := []int{1, 2, 3, 4, 5} // スライスの内容を指定し、配列と同時にスライス作成
	fmt.Println("s1=", s1)
	// スライスに要素を追加
	s2 := append(s1, 6, 7)
	fmt.Println("s2=", s2)
	// スライスにスライスを追加
	s3 := append(s1, s2...)
	fmt.Println("s3=", s3)

	// スライスの要素のコピー
	// スライスの要素を別のスライスにコピーするには「copy」関数を使用する。

	// スライスの作成
	src1 := []int{1, 2}
	dest := []int{97, 98, 99}

	//destへsrc1の内容をすべてコピー(上書き) 戻り値に上書き件数が入る。
	count := copy(dest, src1)
	fmt.Println("copy count=", count)
	fmt.Println(dest)
	fmt.Println()

	src2 := []int{3}
	// destの2番目のindexにsrc2をコピー
	count = copy(dest[2:], src2)
	fmt.Println("copy count=", count)
	fmt.Println(dest)
	fmt.Println()
	// makeを使用したスライスの作成
	// make関数を利用すると、初期化時にデータがなくても1度でスライスが作成できます。
	// string型スライスを作成
	slc1 := make([]string, 5, 10)
	fmt.Println("len=", len(slc1))
	fmt.Println("cap=", cap(slc1))
	fmt.Println()
	// capを省略
	slc2 := make([]string, 5)
	fmt.Println("len=", len(slc2))
	fmt.Println("cap=", cap(slc2))

}
