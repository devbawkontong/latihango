package main

import (
	"fmt"
	"strings"
)


var x string = "Hello World"

func ff(){
	fmt.Println(x)
}

func scan(){
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func main(){
	ff()
	var input string
	fmt.Scanln(&input)

	fmt.Println(strings.Contains("test", "e"))

	m := "latihango"
	fmt.Println(len(m))

	scan()
}
