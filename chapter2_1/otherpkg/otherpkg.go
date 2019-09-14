package otherpkg

// importで他packageにアクセスする
import "github.com/ktr11/go_tutorial1/chapter2_1/somepkg"

func OtherFunc() {
	somepkg.SomeFunc()
	somepkg.SomeVar = 5
}
