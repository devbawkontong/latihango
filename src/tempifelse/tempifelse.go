package main

import "fmt"
import "bufio"
import "os"
//import "strconv"

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Print(text)
	//var point, _ = strconv.ParseFloat(text, 64)
	var float = 100.23
	fmt.Printf("point is %f \n", float)
	if percent := float / 100; percent >= 100 {
		fmt.Printf("%.2f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.2f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.2f%s not bad\n", percent, "%")
	}
}