// Goのラベル
package main

import "fmt"

func main() {
	// break文のラベル
	// 通常のbreak文は、所属するもっとも内側のfor文、switch文、select文から抜け出すが、
	// ラベルを使用すると、任意の文から抜け出すことができる。
FOR_LABEL:
	for i := 0; i < 10; i++ {
		switch {
		case i == 3:
			// for 文からの脱出
			break FOR_LABEL
		default:
			fmt.Println(i)
		}
	}

	// continue文のラベル
	// break文と同様、任意のfor文の実行を中断させることができる。
LABEL1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 && j == 5 {
				// 1番目のforへcontinue
				fmt.Printf("i:%d,j:%d\n", i, j)
				continue LABEL1
			} else if i == 1 && j == 1 {
				// 2番目のforへcontinue
				fmt.Printf("i:%d,j:%d\n", i, j)
				continue
			}
		}
	}

	// goto文でのラベル
	// 関数内のラベルが宣言された場所に移動します。
	for i := 0; i < 10; i++ {
		fmt.Printf("i:%d\n", i)
		if i == 2 {
			// for文の外にあるLABELへ移動
			goto LABEL
		}
	}
LABEL:
}
