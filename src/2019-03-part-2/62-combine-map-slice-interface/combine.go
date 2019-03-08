package main

import "fmt"

func main() {
	var person = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}
	for _, gg := range person {
		fmt.Println(gg["name"], "age is", gg["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	for _, each := range fruits {
		fmt.Println(each)
	}
}
