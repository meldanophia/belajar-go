package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// kalo kita bikin object customer-nya, baru bisa panggil sayHello-nya
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello " + name + ". My name is " + customer.Name)
}

func main() {
	//dibuat dlu si object budi berdasarkan struct customer,
	budi := Customer{"Budi", "Semarang", 35}

	//baru bisa panggil si sayHello-nya
	budi.sayHello("Agus")
}
