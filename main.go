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
	case "observer":
		print("nao implementado ainda")
	default:
		fmt.Print("Método desconhecido")
	}
}
