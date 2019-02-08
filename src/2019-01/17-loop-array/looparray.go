package main

import "fmt"

func main(){
	var buah = [4]string{"pisang", "mangga", "duren", "jeruk"}
	for i:=0;i<len(buah);i++{
		fmt.Printf("Elemen %d : %s\r\n", i, buah[i])
	}
	fmt.Println("------------------")
	for i, b := range buah {
		fmt.Printf("Elemen %d : %s\r\n", i, b)
	}
	fmt.Println("------------------")
	for _, b := range buah {  //index ditampung underscore
		fmt.Println(b)
	}
	fmt.Println("------------------")
	for b := range buah {  //index ditampung ke b, element tidak ada
		fmt.Println(b)
	}
}