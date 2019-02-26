package main

import "fmt"
import "math"

func main() {
	var dia float64 = 5
	var lu, kel = kalkulasi2(dia)

	fmt.Printf("luas : %.2f\n", lu)
	fmt.Printf("keliling : %.2f\n", kel)

}

func kalkulasi(diameter float64) (float64, float64) {
	//hitung luas
	var luas = math.Phi * math.Pow(diameter/2, 2)
	//hitung keliling
	var keliling = math.Phi * diameter
	return luas, keliling
}

//cara lain return-nya
func kalkulasi2(diameter float64) (luas float64, keliling float64) {
	//hitung luas
	luas = math.Phi * math.Pow(diameter/2, 2)
	//hitung keliling
	keliling = math.Phi * diameter
	return
}
