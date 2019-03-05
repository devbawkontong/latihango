package main

import "fmt"

func main() {
	var number = 4
	fmt.Println("before : ", number)
	change(&number, 15)
	fmt.Println("after  : ", number)
}

func change(ori *int, value int) {
	*ori = value
}
