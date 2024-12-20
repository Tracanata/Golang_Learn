package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (hewan Animal) GetName() string {
	return hewan.Name
}

func main() {
	person := Person{Name: "Eko"}
	sayHello(person)

	hewanKucing := Animal{Name: "Kucing"}
	sayHello(hewanKucing)
}
