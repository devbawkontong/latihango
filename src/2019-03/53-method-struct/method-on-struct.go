package main

import "fmt"
import "strings"

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var st = student{"john wick", 22}
	st.sayHello()
	fmt.Println("last name :", st.getNameAt(2))
}
