package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Print(text)
	var point, _ = strconv.ParseFloat(strings.TrimSpace(text), 64)

	fmt.Printf("point is %f \n", point)
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.2f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.2f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.2f%s not bad\n", percent, "%")
	}
}
