package main

import "github.com/googlr/uuid"

func main() {
	// trabalhar como se quisesse imprimir um uuid
	// digitar o import "github.com/googlr/uuid" mas ainda sem dar o go mod tidy, e ele não funcionará
	println(uuid.New().String())
}
