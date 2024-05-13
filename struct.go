package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var melda Customer
	melda.Name = "Melda Nophia"
	melda.Address = "Bekasi"
	melda.Age = 25

	fmt.Println(melda)
}
