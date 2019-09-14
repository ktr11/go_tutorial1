package otherpkg

// importで他packageにアクセスする
import "somepkg"

func OtherFunc() {
	somepkg.SomeFunc()
	somepkg.SomeVar = 5
}
