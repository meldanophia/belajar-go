package main

import "fmt"

func main() {
	// map[type]type
	// map[key-nya tipe apa] value-nya tipe apa

	///// map could be initialized this way
	// var person map[string]string = map [string]string{}
	// person["name"] = "Melda"
	// person["address"] = "Bekasi"

	person := map[string]string{
		"name":    "Melda",
		"address": "Bekasi",
	}

	fmt.Println(person["name"])    //Melda
	fmt.Println(person["address"]) //Bekasi

	book := make(map[string]string)
	book["title"] = "how to be person"
	book["author"] = "melda"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book) //map[author:melda title:how to be person]
}
