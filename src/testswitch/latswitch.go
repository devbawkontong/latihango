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

	var point, _ = strconv.ParseInt(strings.TrimSpace(text), 10, 0)
	fmt.Printf("point is %d\n",point)

	switch point {
	case 10:
		fmt.Println("perfect")
	case 7,8,9:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}