package main

import "fmt"

func main(){
	var fruits=[]string{"mangga", "jeruk", "apel", "jambu"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(aFruits)
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4
	
	var bFruits = fruits[1:4]
	fmt.Println(bFruits)
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3
}