package main

import "fmt"

func sayHello(name string) string {
	return "hai " + name
}

func main() {
	result := sayHello("melda")
	fmt.Println(result)
	fmt.Println(sayHello("nophia"))
}
