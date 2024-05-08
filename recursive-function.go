package main

import "fmt"

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func factorialLoopMain() {
	fmt.Println(factorialLoop(10))
}

func factorialRecursiveMain() {
	fmt.Println(factorialRecursive(10))
}
