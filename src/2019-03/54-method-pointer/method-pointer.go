package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s student) changeName1(name string) {
	fmt.Println("change1 name to :", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("change2 name to :", name)
	s.name = name
}

func main() {
	var st = student{"john rambo", 35}
	fmt.Println("name :", st.name)
	st.changeName1("harry potter")
	fmt.Println("name change 1 :", st.name)
	st.changeName2("james bond")
	fmt.Println("name change 2 :", st.name)
}
