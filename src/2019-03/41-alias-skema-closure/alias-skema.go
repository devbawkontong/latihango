package main

import (
	"fmt"
	"strings"
)

type FilterCallBack func(string) bool

func filter(data []string, callback FilterCallBack) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func main() {
	var data = []string{"golang", "python", "java"}
	var dataContaintGO = filter(data, func(each string) bool {
		return strings.Contains(each, "go")
	})

	fmt.Println("data asli :", data)
	fmt.Println("data ada 'go' :", dataContaintGO)
}
