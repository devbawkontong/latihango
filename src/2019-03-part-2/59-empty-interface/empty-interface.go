package main

import "fmt"

func main() {
	var test interface{}

	test = "test empty interface"
	fmt.Println(test)

	test = []string{"jeruk", "mangga", "pisang"}
	fmt.Println(test)

	test = map[int]string{
		1: "Burger",
		2: "Doughnout",
		3: "Pizza",
	}
	fmt.Println(test)

	test = 12.4
	fmt.Println(test)
}
