// インターフェース
package main

import "fmt"

type Animal interface{
	Cry()
}

type Dog struct {
}

func (d *Dog) Cry() {
	fmt.Println("わんわん")
}

type Cat struct {
}

func (c *Cat) Cry() {
	fmt.Println("にゃんにゃん")
}

func MakeDogCry(d *Dog) {
	fmt.Println("鳴け！")
	d.Cry()
}

func MakeCatCry(c *Cat) {
	fmt.Println("鳴け！")
	c.Cry()
}

func main() {
	dog := new(Dog)
	cat := new(Cat)
	MakeDogCry(dog)
	MakeCatCry(cat)
}
