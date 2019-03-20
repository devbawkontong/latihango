package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}
func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var cube hitung
	cube = &kubus{4}
	fmt.Println("luas:", cube.luas())
	fmt.Println("keliling:", cube.keliling())
	fmt.Println("volume:", cube.volume())
	fmt.Println("sisi:", cube.(*kubus).sisi) //harus pakai * karena kubus diawal adalah pointer
}
