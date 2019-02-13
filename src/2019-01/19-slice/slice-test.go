package main

import "fmt"

func main(){
	var fruits = []string{"apple", "banana", "orange"}
	fmt.Println(fruits[1])

	//new var
	var newFruits = fruits[0:2]
	fmt.Println(newFruits)

	var newFruits2 = fruits[:]
	fmt.Println(newFruits2)

	fmt.Println("---------------------")
	
	fruits[1] = "peach"
	fmt.Println(newFruits)
	fmt.Println(newFruits2)
}