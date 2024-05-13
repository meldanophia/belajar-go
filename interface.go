package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHi(value HasName) {
	fmt.Println("Hello " + value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

// struct animal mengimplementasikan interface hasName
// karena memiliki funcion getName yg sesuai dgn interface hasName
type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Melda"}
	sayHi(person)

	animal := Animal{Name: "Hachi"}
	sayHi(animal)
}
