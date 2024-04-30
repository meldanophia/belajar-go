package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Melda"
	names[1] = "Nophia"

	fmt.Println(names) //[Melda Nophia]

	var angka = [5]int{
		42342,
		21,
		75453,
	}

	fmt.Println(angka) // [42342 21 75453 0 0] , default value array is 0

	var number = [...]int{
		3,
		3013,
		9301,
		4829,
	}

	fmt.Println(number, len(number)) //[3 3013 9301 4829] 4

}
