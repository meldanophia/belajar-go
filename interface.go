package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHi(value HasName) {
	fmt.Println("Hello " + value.getName())
}

func main() {
}
