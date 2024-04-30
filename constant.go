package main

import "fmt"

func main() {
	const firstName = "Melda"
	const lastName = "Nophia"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// cannot assign to firstName (neither addressable nor a map index expression)
	// firstName = "Meldy"
}
