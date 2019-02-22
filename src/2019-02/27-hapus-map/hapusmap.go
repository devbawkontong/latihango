package main

import "fmt"

func main() {
	var deliciousfood = map[int]string{
		1: "Fried Chicken",
		2: "Burger",
		3: "Pizza",
		4: "Dimsum",
	}
	delete(deliciousfood, 2)
	fmt.Println(deliciousfood)
}
