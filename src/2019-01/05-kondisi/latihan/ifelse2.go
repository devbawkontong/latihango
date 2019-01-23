package latihan

import "fmt"

func Show(){
	var point2 = 8840.0
	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s sempurna!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s bagus\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}