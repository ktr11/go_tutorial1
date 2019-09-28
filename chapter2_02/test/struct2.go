package test

// Vector3 3次元ベクトルを定義
type Vector3 struct {
	X int
	Y int
	z int //小文字から始まるメンバーにはパッケージ外からアクセスできない。
}

// SetZ Vector3型のZに値を設定
func SetZ(v *Vector3, zValue int) {
	v.z = zValue
}

// privateな3次元ベクトルを定義
type privateVector3 struct {
	X int
	Y int
	Z int
}

// NewVector privateVector3型の変数を返す
func NewVector() privateVector3 {
	var v privateVector3
	return v
}
