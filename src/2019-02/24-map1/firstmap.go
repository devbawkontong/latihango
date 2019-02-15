package main

import "fmt"

func main() {
	//var chicken map[string]int
	var chicken = map[string]int{}
	chicken["januari"] = 100
	chicken["maret"] = 10

	fmt.Println("januari :", chicken["januari"])

	//menimpa nilai sebelumnya
	chicken["maret"] = 200
	fmt.Println("maret :", chicken["maret"])

	//ambil key yg kosong, nilai akan berupa default
	fmt.Println("mei :", chicken["mei"])
}
