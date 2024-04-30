package main

import "fmt"

func main() {
	names := [...]string{"melda", "nophia", "budi", "eko", "anton"}

	slices1 := names[2:5] // [budi eko anton]
	fmt.Println(slices1)

	slices2 := names[:3] // [melda nophia budi]
	fmt.Println(slices2)

	slices3 := names[:]
	fmt.Println(slices3) //[melda nophia budi eko anton]

}
