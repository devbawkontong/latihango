package main

import "fmt"

func ratarata(numbers ...int) float64 {
	var total int
	for _, angka := range numbers {
		total += angka
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func ratarata2(numbers []int) float64 {
	var total int
	for _, angka := range numbers {
		total += angka
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func main() {
	var numbers = []int{2, 3, 4, 5, 9}
	var rata = ratarata(numbers...)
	var msg = fmt.Sprintf("rata-rata %.2f", rata)
	fmt.Println(msg)

	var rata2 = ratarata2(numbers)
	var msg2 = fmt.Sprintf("rata-rata %.2f", rata2)
	fmt.Println(msg2)
}
