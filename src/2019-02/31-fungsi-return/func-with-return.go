package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomRange(min, max int) int {
	var value = rand.Int() % (max - min + 1)
	return value
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var randomValue int
	randomValue = randomRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomRange(2, 10)
	fmt.Println("random number:", randomValue)
}
