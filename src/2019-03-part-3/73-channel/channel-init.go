package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	var pesan = make(chan string)

	var sayHalo = func(siapa string) {
		var data = fmt.Sprintf("hellow %s", siapa)
		fmt.Println("-->", data)
		pesan <- data
	}
	go sayHalo("iron man")
	go sayHalo("star-lord")
	go sayHalo("thanos")

	var pesan1 = <-pesan
	fmt.Println(pesan1)

	var pesan2 = <-pesan
	fmt.Println(pesan2)

	var pesan3 = <-pesan
	fmt.Println(pesan3)
}
