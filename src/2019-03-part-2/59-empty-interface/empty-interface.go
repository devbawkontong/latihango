package main

import "fmt"

func main() {
	var test interface{}

	test = "test empty interface"
	fmt.Println(test)

	s := test.(string)
	fmt.Println(s)

	//akan cek tipe-nya, jika OK maka akan set nilai sbg float64, jika tidak di-set default nilai float64 (nol)
	f, ok := test.(float64)
	fmt.Println(f, ok)

	//langsung set nilai, akan keluar panic: interface conversion: interface {} is string, not float64
	//f = test.(float64)
	//fmt.Println(f)

	test = []string{"jeruk", "mangga", "pisang"}
	fmt.Println(test)

	test = map[int]string{
		1: "Burger",
		2: "Doughnout",
		3: "Pizza",
	}
	fmt.Println(test)

	test = 12.4
	fmt.Println(test)
}
