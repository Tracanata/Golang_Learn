package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fatra := Man{"Fatra"}
	fatra.married()

	fmt.Println(fatra.Name)
}
