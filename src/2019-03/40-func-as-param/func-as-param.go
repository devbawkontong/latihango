package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContaintO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli :", data)
	fmt.Println("data ada 'o' :", dataContaintO)
	fmt.Println("data ada 5 :", dataLength5)
}
