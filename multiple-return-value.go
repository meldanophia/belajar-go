package main

import "fmt"

func getFullName() (string, string) {
	return "Melda", "Nophia"
}

func main() {
	// the value of multiple return value must be called
	firstName, LastName := getFullName()
	fmt.Println(firstName, LastName) //Melda Nophia

	//if we only need less than multiple return value, then use _
	firstNameV2, _ := getFullName()
	fmt.Println(firstNameV2) //Melda
}
