package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println("nilai32", nilai32) //32768
	fmt.Println("nilai64", nilai64) //32768
	fmt.Println("nilai16", nilai16) //-32768

	var name string = "melda"
	var m uint8 = name[0]
	fmt.Println(m) //109
	var mToString = string(e)
	fmt.Println(mToString) //m
}
