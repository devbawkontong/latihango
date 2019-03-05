package main

import "fmt"

func main() {
	var numA int
	numA = 4
	var numB *int
	numB = &numA

	fmt.Println("numA value :", numA)
	fmt.Println("numA addr :", &numA)
	fmt.Println("numB value :", *numB)
	fmt.Println("numB addr :", numB)

	numA = 8

	fmt.Println("numA value-changes :", numA)
	fmt.Println("numA addr-changes :", &numA)
	fmt.Println("numB value-changes :", *numB)
	fmt.Println("numB addr-changes :", numB)
}
