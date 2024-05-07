package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40)
	fmt.Println(total) //100

	// The sumAll function expects a variadic parameter (denoted by ...int),
	// which means it can accept zero or more int arguments, not a slice.
	// To fix this, you can use the ... operator to expand the slice into individual arguments
	fmt.Println("total in slices", sumAll([]int{1, 2, 3, 4}...))
}
