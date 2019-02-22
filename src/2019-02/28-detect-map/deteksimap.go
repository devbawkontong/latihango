package main

import "fmt"

func main() {
	var meal = map[int]string{1: "soto", 2: "nasi goreng", 3: "bakso", 4: "sate"}
	var value, isExist = meal[5]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item not exist")
	}
}
