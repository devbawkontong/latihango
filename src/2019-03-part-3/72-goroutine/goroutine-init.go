package main

import "fmt"
import "runtime"

func cetak(batas int, pesan string) {
	for i := 0; i < batas; i++ {
		fmt.Println((i + 1), pesan)
	}
}

func main() {
	runtime.GOMAXPROCS(3)

	go cetak(5, "hola")
	cetak(5, "halo")

	var input, input2, input3 string
	fmt.Scanln(&input, &input2, &input3)
	fmt.Println(input)
	fmt.Println(input2)
	fmt.Println(input3)
}
