package main

import "fmt"

func main(){
	var fruits = []string{"apple", "grape", "banana"}

	//tambah data slice
	var buahBaru=append(fruits, "pepaya", "durian")
	fmt.Println(buahBaru)

	var bFruits = fruits[0:2]
	fmt.Println(cap(bFruits))
	fmt.Println(len(bFruits))

	fmt.Println("---------------")
	fmt.Println(fruits)
	fmt.Println(bFruits)

	fmt.Println("---------------")
	var cFruits = append(bFruits, "papaya", "kedondong")
	fmt.Println(fruits) 
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}