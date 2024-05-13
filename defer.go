package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function ")
}

func runApplication() {
	// whether runApplication failed/succeed,
	// function that called using defer,
	//  will always been called, at the end of the task
	defer logging()

	fmt.Println("run application")
}

func main() {
	runApplication()
}
