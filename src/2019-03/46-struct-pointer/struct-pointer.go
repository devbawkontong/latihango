package main

import "fmt"

type student struct {
	nama  string
	grade int
}

func main() {
	var st1 = student{nama: "bruce", grade: 1}
	var st2 *student
	st2 = &st1
	fmt.Println("nama 1:", st1.nama)
	fmt.Println("nama 2:", st2.nama)
	st2.nama = "harry"
	fmt.Println("nama 1:", st1.nama)
	fmt.Println("nama 2:", st2.nama)

	g := student{"joker", 2}
	fmt.Println(g.nama)
	h := &g
	h.grade = 1e6
	fmt.Println(g)
}
