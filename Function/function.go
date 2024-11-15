package main

/*
// -- FUNCTION --
func sayHello() {
	fmt.Println("Hello World")
}
*/

/*
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
*/

/*
// -- FUNCTION RETURN VALUE --
func getHello(name string) string {
	return "Hello " + name
}
*/

/*
// -- FUNCTION MULTIPLE RETURN VALUE --
func getFullName() (string, string) {
	return "jamal", "ubed"
}
*/

/*
// -- Named_Return_Value --
func getCompleteName() (a string, b string, c string) {
	a = "Fatra"
	b = "Candra"
	c = "Ardiwinata"

	return a, b, c
}
*/

/*
// -- Variadic Function --
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
*/

/*
// -- Function Value --
func getGoodBye(nama string) string {
	return "Good Bye " + nama
}
*/

/*
func main() {
	// cara panggil biasa
	sayHello()
	sayHelloTo("Fatra", "Ardiwinata")
	sayHelloTo("Ubed", "Hitler")

	// cara panggil function return
	result := getHello("Hitler")
	fmt.Println(result)
	fmt.Println(getHello("Fatra"))

	// cara panggil function return multiple value
	nameFisrt, nameLast := getFullName()
	fmt.Println("Hello", nameFisrt, "dan", nameLast)

	// Ingore multiple return vaule
	pertama, _ := getFullName()
	fmt.Println("Hello", pertama)

	// Panggil fungsi named_return_value
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	// Panggil fungsi variadic
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	// Panggil fungsi function as value
	goodBye := getGoodBye
	fmt.Println(goodBye("Hayuuu"))
}
*/
