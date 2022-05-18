package main

import (
	"designer-for-golang/test"
	"fmt"
)

func main() {

	text := test.InputiaOTexto()

	switch text {
	case "provider":
		test.TestProvider()
	case "generator":
		test.TestGenerator()
	case "observer":
		test.TesteObserver()
	default:
		fmt.Print("Método desconhecido")
	}
}
