package main

import "fmt"

func main() {
	name := "Nata"

	if name == "Fatra" {
		fmt.Println("Welcome Back Fatra !!!")
	} else if name == "Candra" {
		fmt.Println("Hi candra, Long time no see")
	} else {
		fmt.Println("who are you?, mother fucker !!!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu Panjang", len(name))
	} else {
		fmt.Println("Nama sudah benar")
	}
}
