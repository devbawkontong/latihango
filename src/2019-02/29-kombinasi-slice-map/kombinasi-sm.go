package main

import "fmt"

func main() {
	var cities = []map[string]string{
		{"city": "surabaya", "province": "jawa timur"},
		{"city": "jakarta", "province": "dki jaya"},
		{"city": "medan", "province": "sumatera utara"},
	}
	for _, city := range cities {
		fmt.Println(city["city"], city["province"])
	}
	fmt.Println()
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	for _, dt := range data {
		fmt.Println(dt)
	}
}
