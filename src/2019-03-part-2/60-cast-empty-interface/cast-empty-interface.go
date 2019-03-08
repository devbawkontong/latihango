package main

import "fmt"
import "strings"

func main() {
	var test interface{}

	test = "test empty interface"
	fmt.Println(test)

	test = []string{"jeruk", "mangga", "pisang"}
	fmt.Println(test)
	//akan error arena tipe masih berupa interface
	//fmt.Println(test[1])
	//harus di-cast
	//Teknik casting pada interface disebut dengan type assertions.
	fmt.Println(test.([]string)[1])
	var fruit = strings.Join(test.([]string), ", ")
	fmt.Println("my favorite fruits are:", fruit)

	test = map[int]string{
		1: "Burger",
		2: "Doughnout",
		3: "Pizza",
	}
	fmt.Println(test.(map[int]string)[2])

	test = 12.4
	fmt.Println(test)
	fmt.Println(test.(float64)) //casting
}
