package main

import "github.com/google/uuid"

func main() {
	// trabalhar como se quisesse imprimir um uuid
	// digitar o import "github.com/google/uuid" mas ainda sem dar o go mod tidy, e ele não funcionará
	println(uuid.New().String())
}

// ao executar  go run main.go o resultado é
// no required module provides package github.com/googlr/uuid; to add it:
// go get github.com/googlr/uuid  ou no caso, dar um go mod tidy pra instalar o pacote
