package main

import "fmt"
import "strings"

func main() {
	var names = []string{"john", "wick"}
	printPesan("hallo", names)
}

func printPesan(pesan string, arr []string) {
	var nameStr = strings.Join(arr, " ")
	fmt.Println(pesan, nameStr)
}
