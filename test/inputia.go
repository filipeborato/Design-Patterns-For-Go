package test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputiaOTexto() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite o método do padrão a ser testado")
	fmt.Println("---------------------")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println("---------------------")

	return text
}
