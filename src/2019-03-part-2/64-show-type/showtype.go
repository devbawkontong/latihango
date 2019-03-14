package main

import "fmt"

const Pi = 3.14

func main() {
	m := 65.45 //coba ganti nilainya dan cek return type-nya!
	fmt.Printf("variabel m bertipe %T\n", m)
	fmt.Println("nilai pi", Pi)

	const World = "世界"
	fmt.Println("Hello", World)
}
