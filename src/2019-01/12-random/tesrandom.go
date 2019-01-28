package main

import (
	"fmt" 
	"math/rand"
	"math"
)

func main(){
	fmt.Println("my favorite number :", rand.Intn(2000))
	fmt.Println("Now you have ", math.Sqrt(7), " problem")
}