package main

import "fmt"

func main(){
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
	fmt.Println("-----------------------")
	for _, fruit := range fruits {
		fmt.Printf("elemen : %s\n", fruit)
	}
}