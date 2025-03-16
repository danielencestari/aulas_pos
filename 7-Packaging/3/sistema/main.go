package main

import "github.com/danielencestari/aulas_pos/7-Packaging/3/math"

func main() {
	// agora quero usar o pacote math
	// para isso, preciso importar o pacote

	m := math.NewMath(1, 2)
	println(m.Add())
}

// formas de publicar o pacote math pra usa-lo aqui
// comando go mod edit -replace github.com/danielencestari/aulas_pos/7-Packaging/3/math=../math
