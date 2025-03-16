package math

// type Math struct {
// 	A int
// 	B int
// }

// func (m Math) Add() int {
// 	return m.A + m.B
// }

// caso não queria os dados publicos/exportados como acima

type math struct {
	a int
	b int
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}

func (m *math) Add() int {
	return m.a + m.b
}
