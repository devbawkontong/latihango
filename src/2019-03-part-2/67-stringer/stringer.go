package main

import "fmt"

/*
Interface yang ada dimanapun yaitu Stringer didefinisikan oleh paket fmt.

type Stringer interface {
    String() string
}
Sebuah Stringer adalah suatu tipe yang mendeskripsikan dirinya sendiri sebagai string.
Paket fmt (dan banyak lainnya) menggunakan interface ini untuk mencetak nilai.
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
