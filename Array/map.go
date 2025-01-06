package main

import "fmt"

func main() {
	// CARA LAMA
	// var person map[string]string = map[string]string{}
	// person["name"] = "Fatra"
	// person["address"] = "Cianjur"

	// CARA CEPAT
	person := map[string]string{
		"name":    "Fatra",
		"address": "Cianjur",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
	fmt.Println(person["agama"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)

}
