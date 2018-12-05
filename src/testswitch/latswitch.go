package main

import (
	"strings"
	"os"
	"bufio"
	"strconv"
	"fmt"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')

	var point, _ = strconv.ParseInt(strings.TrimSpace(text), 16, 32)

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}