package main

import (
	"fmt"

	// consigo usar o math pq importei o pacote
	"github.com/danielencestari/aulas_pos/7-Packagin/1/math"
)

// func main() {
// 	m := math.Math{A: 5, B: 6}
// 	fmt.Println(m.Add())
// }

// modo 'privado'
func main() {
	m := math.NewMath(5, -1)
	fmt.Println(m.Add())
}
