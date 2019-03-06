package main

import "fmt"

type employee struct {
	name  string
	grade int
}

func main() {
	var allEmployee = []employee{
		{name: "rocket", grade: 4},
		{name: "groot", grade: 7},
		{name: "drax", grade: 6},
	}

	for _, emp := range allEmployee {
		fmt.Println(emp.name, ":", emp.grade)
	}
}
