package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(4)

	message := make(chan int, 2)
	go func() {
		for {
			i := <-message
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("sending data", i)
		message <- i
	}

	var w string
	fmt.Scan(&w) //buat mastikan semua goroutine terbaca oleh for
}
