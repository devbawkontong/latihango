package main

import "fmt"

func main(){
	var positif uint = 999
	var negatif = -123

	fmt.Printf("Nilai : %d\n", positif)
	fmt.Printf("Nilai : %d\n", negatif)

	var desimal = 2.3419
	fmt.Printf("Nilai des : %.13f\n",desimal)

	var ada = false
	fmt.Printf("ada gak sih? %t\n", ada)

	var moto = "carpe diem"
	fmt.Printf("Moto : %s\n", moto)

	var msg = `I am "Batman" 
I will kill Darkseid`
	fmt.Println(msg)

	//konstanta
	const name = "james bond"
	//name = "jason bourne"  error const cannot be override
	fmt.Println("hola", name)
}