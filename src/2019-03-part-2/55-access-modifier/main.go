package main

import lib2 "2019-03-part-2/55-access-modifier/library"

func main() {
	//lib.SayHello()
	//modified to
	lib2.SayHello("barry ellen")

	//tidak bisa diakses karena modifier-nya private,
	//nama method diawali huruf kecil "introduce"
	//lib.introduce("barry allen")
}
