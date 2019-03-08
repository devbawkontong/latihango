package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var tt interface{} = &person{name: "james bond", age: 40}
	var name = tt.(*person).name
	fmt.Println("name :", name)
}
