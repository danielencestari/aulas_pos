package math

// go mod init github.com/danielencestari/aulas_pos/7-Packagin/3/math
// comando para criar o go.mod do pacote math
// pra criar o go.mod do sistema, rodar o comando go mod init github.com/danielencestari/aulas_pos/7-Packagin/3/sistema

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
