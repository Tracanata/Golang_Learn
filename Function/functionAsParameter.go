package main

// <-- cara biasa -->
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello ", filteredName)
// }

// // Menggunakan Type Declaration
// type Filter func(string) string

// func sayHelloWithFilter(name string, filter Filter) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello ", filteredName)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	sayHelloWithFilter("Eko", spamFilter)

// 	filter := spamFilter
// 	sayHelloWithFilter("Anjing", filter)
// }
