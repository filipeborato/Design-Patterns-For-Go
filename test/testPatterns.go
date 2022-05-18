package test

import (
	"designer-for-golang/factory"
	"designer-for-golang/generator"
	"designer-for-golang/observer"
	"designer-for-golang/provider"
	"fmt"
	"log"
	"os"
	"reflect"
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

func TestFactory() {

	env1 := "production"
	env2 := "development"

	db1, fs1 := SetupConstructorsFactory(env1)
	db2, fs2 := SetupConstructorsFactory(env2)

	// db1 := factory.DatabaseFactory(env1)
	// db2 := factory.DatabaseFactory(env2)

	db1.PutData("test", "this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite")
	fmt.Println(db2.GetData("test"))

	fs1.CreateFile("../example/testntfs.txt")
	fmt.Println(fs1.FindFile("../example/testntfs.txt"))

	fs2.CreateFile("../example/testext4.txt")
	fmt.Println(fs2.FindFile("../example/testext4.txt"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())

	fmt.Println(reflect.TypeOf(fs1).Name())
	fmt.Println(reflect.TypeOf(&fs1).Elem())
	fmt.Println(reflect.TypeOf(fs2).Name())
	fmt.Println(reflect.TypeOf(&fs2).Elem())

}

func SetupConstructorsFactory(env string) (factory.Database, factory.FileSystem) {
	fs := factory.AbstractFactory("filesystem")
	db := factory.AbstractFactory("database")

	return db(env).(factory.Database),
		fs(env).(factory.FileSystem)
}
