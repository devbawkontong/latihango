package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	fmt.Println("-----------------")

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pinnaple"}
	n = copy(dst, src)
	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	//fungsi lain make
	b := make([]int, 5, 5) // len(b)=0, cap(b)=5
	fmt.Println(b)
}
