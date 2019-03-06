package main

import "fmt"

type student struct {
	nama  string
	grade int
}

func main() {
	var st student
	st.nama = "hendy"
	st.grade = 6

	fmt.Println("nama  :", st.nama)
	st.grade = 7
	fmt.Println("grade :", st.grade)

	var st2 = student{nama: "test"}
	fmt.Println("nama  :", st2.nama)
	fmt.Println("grade  :", st2.grade)
}
