package main

import (
	"fmt"
	"strings"
)

/*
// Interface
type Hitung interface {
	luas() float64
	keliling() float64
}

func bangunDatar(hitung Hitung) {
	fmt.Println("Luas : ", hitung.luas())
	fmt.Println("Keliling : ", hitung.keliling())
}

func bangunRuang(hitung Hitung) {
	fmt.Println("Luas Permukaan : ", hitung.luas())
	fmt.Println("Keliling : ", hitung.keliling())
}

type kubus struct {
	sisi float64
}

func (k kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k kubus) luas() float64 {
	return 6 * math.Pow(k.sisi, 2)
}

func (k kubus) keliling() float64 {
	return 12 * k.sisi
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

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

// End Interface
*/

type person struct {
	name string
	age  int
}

func main() {
	/*
		// <--- Interface --->
		person := Person{Name: "Eko"}
		sayHello(person)

		hewanKucing := Animal{Name: "Kucing"}
		sayHello(hewanKucing)

		fmt.Println("====== LINGKARAN =====")
		hitungLingkaran := lingkaran{diameter: 14.0}
		fmt.Println("jari jari : ", hitungLingkaran.jariJari())
		bangunDatar(hitungLingkaran)

		fmt.Println("==== KUBUS ====")
		hitungKubus := kubus{sisi: 10}
		fmt.Println("Volume : ", hitungKubus.volume())
		bangunRuang(hitungKubus)

	*/

	var secret interface{}

	secret = "tidak ada tahu"
	fmt.Println(secret)

	secret = []string{"SIOMAY", "Koll", "Telur"}
	var partBatagor = strings.Join(secret.([]string), ", ")
	fmt.Println(partBatagor, "is my favorite part in batagor")

	secret = 3.14
	fmt.Println(secret)

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "Multiplied by 10 is :", number)

	var secret2 interface{} = &person{name: "fatra", age: 22}
	var name = secret2.(*person).name
	println(name)

	var person = []map[string]interface{}{
		{"name": "Fatra", "age": "22"},
		{"name": "Candra", "age": "22"},
		{"name": "Ardiwinata", "age": "23"},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is ", each["age"])
	}
}
