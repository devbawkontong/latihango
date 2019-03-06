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
}
