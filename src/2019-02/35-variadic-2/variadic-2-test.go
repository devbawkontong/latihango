package main

import (
	"fmt"
	"strings"
)

func hobiku(name string, hobi ...string) {
	var hobies = strings.Join(hobi, ",")

	fmt.Println("my name is :", name)
	fmt.Println("my hobbies are:", hobies)
}

func main() {
	hobiku("hendy", "eating", "sleeping", "browsing")
}
