package main

import "fmt"

func main(){
	rang := 5
	for i:=1;i<=rang;i++ {
		for h:=0;h < rang-i;h++{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}