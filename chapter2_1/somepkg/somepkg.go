package somepkg

// SomeVar int
var SomeVar int
var someVar2 int

// SomeFunc publicな関数
func SomeFunc() {
	SomeVar = 10
	someVar2 = 5
}

func someFunc2() {
	SomeFunc()
}

package otherpkg

// importで他packageにアクセスする
import "somepkg"

func OtherFunc() {
	somepkg.SomeFunc()
	somepkg.SomeVar = 5
}
