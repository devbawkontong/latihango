package main

import "fmt"

func main(){
  var fruits = [...]string{"apple", "banana", "grape", "orange", "melon"}
  fmt.Println("data array \t:", fruits)
  fmt.Println("jumlah elemen \t:", len(fruits))
}
