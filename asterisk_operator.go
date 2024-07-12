package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"

	fmt.Println(address1) // {Subang Jawa Barat Indonesia}
	fmt.Println(address2) // {Bandung Jawa Barat Indonesia}

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // {Jakarta DKI Jakarta Indonesia}
	fmt.Println(address2) // &{Jakarta DKI Jakarta Indonesia}

}
