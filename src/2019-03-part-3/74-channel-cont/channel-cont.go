package main222

import "fmt"
import "runtime"

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	runtime.GOMAXPROCS(2)

	var msg = make(chan string)
	var slice = []string{"wick", "hunt", "bourne"}

	for idx, tiap := range slice {
		fmt.Println(idx)
		go func(siapa string) {
			var data = fmt.Sprintf("hello %v", siapa)
			msg <- data
		}(tiap)
	}

	for i := 0; i < 3; i++ {
		printMessage(msg)
	}
}
