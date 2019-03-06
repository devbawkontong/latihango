package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var s1 = struct {
		person
		grade int
	}{}

	var s2 = struct {
		person
		grade int
	}{
		person: person{name: "groot", age: 5},
		grade:  5,
	}

	s1.person = person{name: "rocket", age: 15}
	s1.grade = 6
	fmt.Println("name1 :", s1.name)
	fmt.Println("age1 :", s1.age)
	fmt.Println("grade1 :", s1.grade)

	fmt.Println("name2 :", s2.name)
	fmt.Println("age2 :", s2.age)
	fmt.Println("grade2 :", s2.grade)
}
