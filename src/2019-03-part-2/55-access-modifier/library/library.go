package lib2

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello world")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("my name is", name)
}
