package main

import "fmt"

func main() {
	var deliciousfood = map[int]string{
		1: "Fried Chicken",
		2: "Burger",
		3: "Pizza",
		4: "Dimsum",
	}

	//iterasi map
	for key, val := range deliciousfood {
		fmt.Println(key, ":", val)
	}
	fmt.Println()
	//iterasi hanya val
	for _, val := range deliciousfood {
		fmt.Println(":", val)
	}
	fmt.Println()
	//iterasi hanya key
	for key, _ := range deliciousfood {
		fmt.Println(key, ":")
	}
}
