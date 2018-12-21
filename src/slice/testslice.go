package main

import "fmt"

func main(){
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0], fruits[2])

	var newFruits = fruits[:]
	fmt.Println(newFruits)
}