package otherpkg

// importで他packageにアクセスする
// packageに別名をつけることもできる。
import some "github.com/ktr11/go_tutorial1/chapter2_1/somepkg"

// OtherFunc publicな関数
func OtherFunc() {
	some.SomeFunc()
	some.SomeVar = 5
}
