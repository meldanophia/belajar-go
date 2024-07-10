package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	fmt.Println(result) // OK

	// implementing type assertions using .(data-type)
	resultString := result.(string)
	fmt.Println(resultString) // OK

	// invalid operation: resultString (variable of type string) is not an interface
	// resultInt := resultString.(int) // panic
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
