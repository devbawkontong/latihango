package main

import "fmt"
import "runtime"
import "time"

func main() {
	point := 8

	switch point {
	case 10, 9, 8:
		fmt.Println("Bleketaket")
	case 7:
		fmt.Println("Apik")
	case 6, 5:
		fmt.Println("lumayan lah")
	default:
		fmt.Println("sing penting ana nilai-ne")
	}

	point = 4
	switch {
	case point >= 8:
		fmt.Println("TOP BGT")
	case (point >= 3) && (point <= 7):
		fmt.Println("LUMAYAN")
		fallthrough
	case point < 5:
		fmt.Println("SINAU MANING")
	}

	fmt.Print("Go berjalan pada ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Print("Kapan hari Sabtu?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Sekarang.")
	case today + 1:
		fmt.Println("Besok.")
	case today + 2:
		fmt.Println("Dua hari lagi.")
	default:
		fmt.Println("Masih jauh.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Selamat pagi!")
	case t.Hour() < 17:
		fmt.Println("Selamat sore.")
	default:
		fmt.Println("Selamat malam.")
	}
}
