package main

import "fmt"

// func Ups() interface{} {
//return 1
//return true
// }

func Ups() any {
	// return 1
	// return true
	return "ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
