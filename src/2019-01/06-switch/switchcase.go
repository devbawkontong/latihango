package main

import "fmt"

func main(){
	point := 8

	switch point{
	case 10,9,8 :
		fmt.Println("Bleketaket")
	case 7 :
		fmt.Println("Apik")	
	case 6,5 :
		fmt.Println("lumayan lah")		
	default :
		fmt.Println("sing penting ana nilai-ne")	
	}


	point = 4
	switch{
	case point >= 8 :
		fmt.Println("TOP BGT")	
	case (point >= 3) && (point <= 7) :
		fmt.Println("LUMAYAN")	
		fallthrough
	case point < 5 :
		fmt.Println("SINAU MANING")
	}
}