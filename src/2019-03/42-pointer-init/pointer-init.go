package main

import "fmt"

func main() {
	var numA int = 4
	var numB *int = &numA

	fmt.Println("numA value :", numA)
	fmt.Println("numA addr :", &numA)
	fmt.Println("numB value :", *numB)
	fmt.Println("numB addr :", numB)
}
