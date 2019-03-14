package main

import "fmt"

func main() {
	//deklarasi var langsung
	var firstName string = "forrest"

	//deklarasi var 2 baris
	var lastName string
	lastName = "gump"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	//deklarasi var tanpa type
	lastNameNew := "bubba"
	lastNameNew = "capt dan" //berikutnya set nilai tanpa :
	two := 2
	two = 10
	fmt.Println("halo", firstName, lastNameNew+"!", two)

	//deklarasi multi var
	var satu, dua, tiga int
	satu, dua, tiga = 1, 2, 3
	fmt.Println(satu, dua, tiga)

	//deklarasi multi var cara langsung
	empat, lima, enam := 4, 5, 6
	fmt.Printf("%d%d%d\n", empat, lima, enam)

	//deklarasi multi var beda-beda tipe
	one, isFriday, duakomadua, say := 1, true, 2.2, "hello"
	fmt.Printf("%d %t %f %s", one, isFriday, duakomadua, say)

	//underscore a.k.a blackhole
	_ = "test..."

	//pointer
	nama := new(string)
	//nama = "hendy"
	fmt.Println(*nama)
}
