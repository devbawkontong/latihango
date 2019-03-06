package library

import "fmt"

/*
type Student struct {
	Name string
	Age  int
}
*/
var Student = struct {
	Name string
	Age  int
}{}

func init() {
	Student.Name = "Star-Lord"
	Student.Age = 35
	fmt.Println("automatic generate data")
}
