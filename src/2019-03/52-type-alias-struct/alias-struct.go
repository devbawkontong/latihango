package main

import "fmt"

type orang struct {
	nama string
	umur int
}

type wong = orang

func main() {
	var s1 = orang{"bubu", 30}
	fmt.Println(s1)
	var s2 = wong{"mimi", 25}
	fmt.Println(s2)
}
