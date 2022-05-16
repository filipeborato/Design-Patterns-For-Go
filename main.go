package main

import (
	"bufio"
	"designer-for-golang/provider"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite o método do padrão a ser testado")
	fmt.Println("---------------------")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println("---------------------")

	if strings.Compare("provider", text) == 0 {
		testeProvider()
	}
}

func testeProvider() {
	f := provider.Wrapcache(provider.Pi, &sync.Map{})
	g := provider.Wraplogger(f, log.New(os.Stdout, "Test ", 1))

	g(100000)
	g(20000)
	g(100000)

	f = provider.Wrapcache(provider.Divide, &sync.Map{})
	g = provider.Wraplogger(f, log.New(os.Stdout, "Divide ", 1))

	g(10000)
	g(2000)
	g(10)
	g(10000)
}
