package main

import (
	"fmt"
)

func main() {
	// -- Konversi Angka --
	var nilai32 int32 = 32767
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// -- Konversi String --

	var name = "Fatra Ardiwinata"
	var huruf = name[0]
	var hurufString = string(huruf)

	fmt.Println(name)
	fmt.Println(huruf)
	fmt.Println(hurufString)
}
