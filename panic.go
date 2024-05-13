package main

import "fmt"

func endApp() {
	fmt.Println("app end")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}

func main() {
	runApp(true)
}
