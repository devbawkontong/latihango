package main

import "fmt"

func main(){
	var value = (((2 + 6) % 3) * 4 - 2) / 3
	var value2 = 0
	//var isEqual = (value == 3)
	var isEqual bool
	if value != 3 || value == 2 && value2 == 4{
		isEqual = true
	}
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}