package main

import (
	"fmt"
)

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke : ", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke : ", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Fatra", "Candra", "Ardiwinata"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, " = ", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	// for i := 0; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}

	// 	fmt.Println("Perulangan ke: ", i)
	// }

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke: ", i)
	}

}
