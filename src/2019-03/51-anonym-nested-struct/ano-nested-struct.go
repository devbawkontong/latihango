package main

import "fmt"

type student struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

func main() {
	var pp student
	pp.person.name = "sule"
	pp.person.age = 40
	pp.grade = 4
	pp.hobbies = []string{"eat", "sleep"}
	fmt.Println(pp.hobbies[0])
}
