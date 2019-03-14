package main

import "fmt"

func main() {
	nilai := 7.999
	switch {
	case nilai >= 8:
		if nilai == 10 {
			fmt.Println("SUPER")
		} else if nilai >= 9 && nilai < 10 {
			fmt.Println("EXCELLENC")
		} else {
			fmt.Println("GREAT")
		}
	case nilai >= 7:
		fmt.Println("GOOD")
	default:
		fmt.Println("TRY AGAIN")
	}
}
