package main

import "fmt"

func main() {
	/*
		Perintah defer menunda eksekusi dari sebuah fungsi sampai fungsi yang melingkupinya selesai.

		Argumen untuk pemanggilan defer dievaluasi langsung, tapi pemanggilan fungsi tidak dieksekusi
		sampai fungsi yang melingkupinya selesai.
	*/
	defer fmt.Println("world") //world jadi paling akhir dikeluarkan krn dibawahnya ada defer lain

	fmt.Println("hello")
	fmt.Println("cruel")

	//contoh lain
	fmt.Println("counting")
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
