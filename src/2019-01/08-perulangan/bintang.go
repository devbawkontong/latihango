package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	/*
		loop forever

		for {
		}

	*/

	//latihan bintang
	rang := 5
	for i := 1; i <= rang; i++ {
		for h := 0; h < rang-i; h++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
