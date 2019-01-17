package main

import (
	"fmt"
	"strings"
)

var x = "Hello World"

func ff(){
	fmt.Println(x)
}

func main(){
	ff()
	var input string
	fmt.Scanln(&input)

	fmt.Println(strings.Contains("test", "e"))
}