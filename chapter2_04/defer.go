// 遅延実行
package main

import (
	"fmt"
	"os"
)

func main() {
	// defer を使用することで対象の処理が関数内の一番最後に実行される。
	// 複数あれば、最初にdeferしたものが最後になる。
	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")
	fmt.Println("D")
	fmt.Println("-------------------------")
	sub()
	fmt.Println()
	fmt.Println("-------------------------")
	// os.Exit()するとdeferした処理も実行されずに終わる
	os.Exit(0)
}

func sub() {
	file, err := os.Open("hoge.txt")
	if err != nil {
		fmt.Println("File Open error: ", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println("File Read error: ", err)
			return
		}
		fmt.Printf("%d:\t", n)
		fmt.Print(buf[:n])
	}
}
