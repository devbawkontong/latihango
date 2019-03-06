package main

import "2019-03-part-2/56-acc-mod-struct/library"
import f "fmt"

func main() {
	//var p = library.Student{Name: "Doel", Age: 19}
	//f.Println(p.Name)
	//f.Println(p.Age)
	f.Println(library.Student.Name)
	f.Println(library.Student.Age)
	sayHi("golang..")
}
