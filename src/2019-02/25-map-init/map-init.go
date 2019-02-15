package main

import "fmt"

func main() {
	var junkfood = map[int]string{
		1: "Burger",
		2: "Dognout",
		3: "Pizza",
	}
	fmt.Println(junkfood[3])

	var drink = make(map[int]string)
	drink[1] = "coca cola"
	drink[2] = "fanta"
	drink[3] = "sprite"
	drink[4] = "beer"
	fmt.Println("I love", drink[3])
}
