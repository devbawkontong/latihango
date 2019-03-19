package main

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	grade int
	person
}

func main() {
	var p1 = person{age: 20, name: "jojon"}
	var e1 = employee{grade: 3, person: p1}
	//p1.name = "jojon"
	//p1.age = 20
	//p1.grade = 3
	fmt.Println(e1.name)
	fmt.Println(e1.person.name)
}
