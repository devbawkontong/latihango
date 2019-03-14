package main

import "fmt"

func main() {
	var numA int
	numA = 4
	var numB *int
	numB = &numA

	fmt.Println("numA value :", numA)
	fmt.Println("numA addr :", &numA)
	fmt.Println("numB value :", *numB)
	fmt.Println("numB addr :", numB)

	numA = 5
	fmt.Println("numA value :", numA)
	fmt.Println("numB value :", *numB)

	i, j := 42, 2701
	p := &i         // menunjuk ke i
	fmt.Println(*p) // baca i lewat pointer
	*p = 21         // set i lewat pointer
	fmt.Println(i)  // lihat nilai terbaru dari i
	p = &j          // p menunjuk ke j
	*p = *p / 37    // bagi nilai j lewat pointer
	fmt.Println(j)  // lihat nilai terbaru dari j
}
