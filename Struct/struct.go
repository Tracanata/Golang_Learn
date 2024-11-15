package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

// Membuat method dari struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello,", name, "My name is ", customer.Name)
}

func main() {
	var fatra Customer
	fmt.Println(fatra)

	fatra.Name = "Fatra Candra Ardiwinata,"
	fatra.Address = "Cianjur,"
	fatra.age = 23
	fmt.Println(fatra)

	ubed := Customer{
		Name:    "Ubed bin jaelani",
		Address: "Jakarta",
		age:     40,
	}
	fmt.Println(ubed)

	jay := Customer{"Jay", "Cianjur", 24}
	fmt.Println(jay)

	jay.sayHello("Ajay")
}
