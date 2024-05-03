package main

import (
	"fmt"
)

func main() {
	names := [...]string{"melda", "nophia", "budi", "eko", "anton"}

	slices1 := names[2:5] // [budi eko anton]
	fmt.Println(slices1)

	slices2 := names[:3] // [melda nophia budi]
	fmt.Println(slices2)

	slices3 := names[:]
	fmt.Println(slices3) //[melda nophia budi eko anton]

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1) //[Sabtu Minggu]

	//replace
	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(daysSlice1) //[Sabtu Baru Minggu Baru]

	//append
	daysSlice2 := append(daysSlice1, "Libur Baru")
	fmt.Println(daysSlice2) //[Sabtu Baru Minggu Baru Libur Baru]

	//how to do slices
	newSlices := make([]string, 2, 5)
	newSlices[0] = "cape gua"
	newSlices[1] = "komeng"
	fmt.Println(newSlices)      //[cape gua komeng]
	fmt.Println(len(newSlices)) //2
	fmt.Println(cap(newSlices)) //5

	newSlices2 := append(newSlices, "budi")
	fmt.Println(newSlices2)      //[cape gua komeng budi]
	fmt.Println(len(newSlices2)) //3
	fmt.Println(cap(newSlices2)) //5

	newSlices2[0] = "kerja"
	fmt.Println(newSlices2) //[kerja komeng budi]
	fmt.Println(newSlices)  //[kerja komeng]

	//copy slices
	fromSlices := days[:]
	toSlice := make([]string, len(fromSlices), cap(fromSlices))

	copy(toSlice, fromSlices)

	fmt.Println(fromSlices) //[Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
	fmt.Println(toSlice)    //[Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	//diff between array and slices
	iniArray := [...]int{1, 2, 3, 4} //using ... inside []
	iniSlices := []int{1, 2, 3, 4}   // using [] only, doesn't need to determine the len

	fmt.Println(iniArray)
	fmt.Println(iniSlices)

}
