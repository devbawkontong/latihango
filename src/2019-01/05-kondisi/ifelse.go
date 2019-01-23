package main

import (
	"fmt"
	"2019-01/05-kondisi/latihan"
)

func main(){
	var point = 9
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	fmt.Println()

	latihan.Show()
}