package main

import (
	"fmt"
	"strconv"
	"strings"
)

var x string = "Hello World"

func ff() {
	fmt.Println(x)
}

func scan() {
	fmt.Print("Enter a number: ")
	var inputStr string
	fmt.Scanf("%s", &inputStr)
	fmt.Println("nilai string:", inputStr)
	if inputf, err := strconv.ParseFloat(inputStr, 10); err != nil {
		fmt.Println(inputf)
		fmt.Println("error not number")
	} else {
		output := inputf * 2
		fmt.Println(output)
	}
}

func main() {
	ff()
	var input string
	fmt.Scanln(&input)

	fmt.Println(strings.Contains(input, "e"))

	m := "latihango"
	fmt.Println(len(m))

	scan()
}
