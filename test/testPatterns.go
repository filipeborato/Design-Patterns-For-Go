package test

import (
	"designer-for-golang/generator"
	"designer-for-golang/observer"
	"designer-for-golang/provider"
	"fmt"
	"log"
	"os"
	"sync"
)

func TestProvider() {
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

func TestGenerator() {
	for x := range generator.Fib(10000000) {
		fmt.Println(x)
	}
}

func TesteObserver() {
	observer.Test()
}
