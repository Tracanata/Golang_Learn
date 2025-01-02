package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Fatra"
	name[1] = "Candra"
	name[2] = "Ardiwinata"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// Membuat array secara langsung

	var values = [...]int{
		90,
		80,
		95,
		100,
		101,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
