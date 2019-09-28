package somepkg

// SomeVar int
var SomeVar int
var someVar2 int

// SomeFunc 名前の一文字目が大文字publicな関数
func SomeFunc() {
	SomeVar = 10
	someVar2 = 5
}

// someFunc2 名前の一文字目が小文字privateな関数
func someFunc2() {
	SomeFunc()
}
