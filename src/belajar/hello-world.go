package main

import "fmt"

func main() {
	var firstName = "john"
	lastName, _ := "wick", "test"
	fmt.Println("halo", firstName, lastName)

	name := new(string)
	fmt.Println(name) // 0x20818a220
	fmt.Println(*name) // ""


	var message = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`
	message = "testing ajah"
	fmt.Println(message)
}