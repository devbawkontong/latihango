package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	p := fmt.Println
	/*
		why referece to 2006-01-02??
		It is just the numbers 1 2 3 4 5 6 7
		1: month (January, Jan, 01, etc)
		2: day
		3: hour (15 is 3pm on a 24 hour clock)
		4: minute
		5: second
		6: year (2006)
		7: timezone (GMT-7 is MST)
	*/
	p(time.Now().Format("2006-01-02"))
	p(time.Now().Format("2006-01-02 15:04:05"))
	p(time.Now().Format("15:04 PM"))

	input := "2017-08-31"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t) // 2017-08-31 00:00:00 +0000 UTC
	fmt.Println(t.Format("02-Jan-2006"))

	//parsing date
	var tgl string = "2018-10-16"
	taim, _ := time.Parse(layout, tgl)
	p(taim)
}
