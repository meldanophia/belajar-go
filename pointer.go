package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}

	// address2 := address1
	// address2.City = "Bandung"

	// fmt.Println(address1) // {Subang Jawa Barat Indonesia}
	// fmt.Println(address2) // {Bandung Jawa Barat Indonesia}

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"

	fmt.Println(address1) // {Bandung Jawa Barat Indonesia}
	fmt.Println(address2) // &{Bandung Jawa Barat Indonesia}

}
