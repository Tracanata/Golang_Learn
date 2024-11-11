package main

import "fmt"

func main() {
	var name string

	name = "Fatra Candra"
	fmt.Println(name)

	name = "Fatra Candra Ardiwinata"
	fmt.Println(name)

	// atau bisa seperti ini

	var name1 = "Fatra Candra Ardiwinata"
	fmt.Println(name1)

	name1 = "Fatra Candra"
	fmt.Println(name1)

	// atau bisa seperti ini

	name2 := "Fatra Candra Ardiwinata"
	fmt.Println(name2)

	name2 = "Fatra Candra"
	fmt.Println(name2)

	// -- Multiple Variabel --

	var (
		firstName  = "Fatra"
		middleName = "Candra"
		lastName   = "Ardiwinata"
	)

	fmt.Println(firstName, middleName, lastName)
}
