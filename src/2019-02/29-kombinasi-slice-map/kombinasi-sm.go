package main

import "fmt"

func main() {
	var cities = []map[string]string{
		{"city": "surabaya", "province": "jawa timur"},
		{"city": "jakarta", "province": "dki jaya"},
		{"city": "medan"},
	}
	for _, city := range cities {
		fmt.Println(city["province"])
	}

}
