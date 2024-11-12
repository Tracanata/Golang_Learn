package main

import "fmt"

// <-- Tanpa menggunakan recursive function -->
// func factorialLoop(value int) int {
// 	result := 1
// 	for i := value; i > 0; i-- {
// 		result *= i
// 	}

// 	return result
// }

// Menggunakan recursive function
func factorialLoop(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
}
