package main

import "fmt"

func main(){
	var names [5]string
	names[0] = "batman"
	names[1] = "superman"
	names[2] = "wonder-woman"
	names[3] = "flash"
	names[4] = "aquaman"

	fmt.Println(names[0], names[1], names[2], names[3], names[4])

	var buah = [4]string{"pisang", "mangga", "duren", "jeruk"}
	fmt.Println(buah)
	fmt.Println("jumlah :\t",len(buah))

	var kota = [3]string{
		"jakarta",
		"surabaya",
		"medan",
	}
	fmt.Println(kota)

	var numbers = [...]int{3,4,5,6}
	fmt.Println(numbers)
	fmt.Println("size numbers : ",len(numbers))
}