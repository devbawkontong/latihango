package main

import "fmt"

func divideBy(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}
	var result = m / n
	fmt.Printf("%d / %d = %d\n", m, n, result)
}

func main() {
	divideBy(4, 0)
}
