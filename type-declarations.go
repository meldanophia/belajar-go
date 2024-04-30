package main

import (
	"fmt"
	// "reflect"
)

func main() {
	type noKTP int

	var ktp noKTP = 394827942
	fmt.Println(ktp)            //394827942
	fmt.Println(noKTP(8923478)) //8923478

	// fmt.Println(reflect.TypeOf(nama))
}
