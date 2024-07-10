package main

import "fmt"

// ini gak bisa dipakai nil-nya, karena return string pada function contoh,
//  bertentangan dengan ketentuan nil
// func contoh (name string) string {
// 	if name == "" {
// 		return nil // error
// 	} else {
// 		return name
// 	}
// }

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// data := NewMap("")      // Data map masih kosong
	data := NewMap("melda") // melda

	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
