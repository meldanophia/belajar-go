package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Ruth"
	middleName = "Irawati"
	lastName = "Setiadi"
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}
