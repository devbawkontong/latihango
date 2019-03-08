package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 23
	var reflexValue = reflect.ValueOf(num)

	fmt.Println("tipe var num :", reflect.TypeOf(num))
	fmt.Println(reflexValue)
	if reflexValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflexValue.Int())
	}
}
