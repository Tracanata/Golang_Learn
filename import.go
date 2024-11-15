package main

import (
	"fmt"
	database "project-pertama/Database"
	"project-pertama/helper"
	_ "project-pertama/internal"
)

func main() {
	result := helper.SayHello("Fatra")
	fmt.Println(result)
	fmt.Println(helper.Application)
	// fmt.Println(helper.sayGoodBye("Fatra")) //tidak bisa diakses
	// fmt.Println(helper.version)             //tidak bisa diakses

	fmt.Println(database.GetDatabase())
}
