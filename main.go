package main

// import (
// 	"fmt"
// )

// func main() {
// 	// fmt.Println("baris ini tidak akan di eksekusi")

// 	// --VARIABEL--
// 	// Deklarasi variabel beserta tipe data (manifest type)
// 	// var firstName string = "Fatra"

// 	// var lastName string
// 	// lastName = "Ardiwinata"

// 	// Deklarasi variabel tanpa tipe data (type inference)
// 	// var firstName = "Fatra"
// 	// lastName := "Ardiwinata"

// 	// fmt.Println("My Name is :", firstName, lastName)
// 	// fmt.Printf("Halo %s %s %s !\n", firstName, lastName)

// 	// Deklarasi multi variabel
// 	// var firstName, middleName, lastName string = "Fatra", "Candra", "Ardiwinata"

// 	// OR WE CAN USE LIKE THIS
// 	// firstName, middleName, lastName := "Fatra", "Candra", "Ardiwinata"

// 	// dapat digunakan juga untuk varibel yang berbeda tipe data
// 	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hi!"

// 	// fmt.Println("My Name is :", firstName, middleName, lastName)
// 	// fmt.Println("execution : ", one, isFriday, twoPointTwo, say)

// 	// fmt.Println("Hello Wolrd! How Are You Today ?")

// 	// -- TIPE DATA --
// 	// Tipe Data Numerik Non Decimal
// 	// var positiveNumber uint8 = 89
// 	// var negativeNumber = -12345678

// 	// fmt.Printf("Bilangan Positif: %d\n", positiveNumber)
// 	// fmt.Printf("Bilangan Negatif: %d\n", negativeNumber)

// 	// Tipe Data Numerik Decimal
// 	// var decimalNumber = 2.62

// 	// fmt.Printf("Bilangan Decimal: %f\n", decimalNumber)
// 	// fmt.Printf("Bilangan Decimal: %.3f\n", decimalNumber)

// 	// Tipe Data Boolean
// 	// var exist = true
// 	// var notExist = false

// 	// fmt.Printf("if Exist ? it's %t\n", exist)
// 	// fmt.Printf("if not Exist ? it's %t\n", notExist)

// 	// Tipe Data String
// 	// var message = "Hello world"
// 	// var message2 = `Hello i'm "Fatra Candra Ardiwinata", let's go to learn about "Golang"`

// 	// fmt.Printf("message : %s\n", message)
// 	// fmt.Printf("message : %s\n", message2)

// 	// -- KONSTANTA --
// 	// fmt.Println("john", "wick")
// 	// fmt.Print("john ", "wick\n")

// 	// Multiple Konstanta
// 	// const (
// 	// 	today string = "Jumat"
// 	// 	sekarang
// 	// 	isToday2 = true
// 	// )
// 	// fmt.Println("Cetak : ", today, sekarang, isToday2)

// 	// -- OPERATOR --
// 	// var value = (((2+6)%3)*4 - 2) / 3
// 	// var isEqual = (value == 2)

// 	// fmt.Println("Result : ", value, isEqual)

// 	// var left = false
// 	// var right = true

// 	// var leftAndRight = left && right
// 	// var leftOrRight = left || right
// 	// var leftRevers = !left

// 	// fmt.Println("result for AND, OR, and Revers is this :", leftAndRight, leftOrRight, leftRevers)

// 	// -- SELEKSI KONDISI --
// 	// var point = 7
// 	// var point = 8840.0

// 	// kondisi dengan if, else if, and else

// 	// if point == 10 {
// 	// 	fmt.Println("Lulus dengan nilai sempurna")
// 	// } else if point > 5 {
// 	// 	fmt.Println("Lulus")
// 	// } else if point == 4 {
// 	// 	fmt.Println("Hampir Lulus")
// 	// } else {
// 	// 	fmt.Printf("Tidak Lulus. Nilai anda %d\n", point)
// 	// }

// 	// Variabel Temporary pada IF - Else
// 	// if percent := point / 100; percent >= 100 {
// 	// 	fmt.Printf("%.1f%s perfect!\n", percent, "%")
// 	// } else if percent >= 70 {
// 	// 	fmt.Printf("%.1f%s good \n", percent, "%")
// 	// } else {
// 	// 	fmt.Printf("%.1f%s not bad \n", percent, "%")
// 	// }

// 	// Seleksi dengan switch case
// 	// switch point {
// 	// case 8:
// 	// 	fmt.Println("Perfect")
// 	// case 7, 6, 5, 4:
// 	// 	fmt.Println("Awesome")
// 	// default:
// 	// 	{
// 	// 		fmt.Println("Not Bad")
// 	// 		fmt.Println("You can be better!")
// 	// 	}
// 	// }

// 	// penulisan switch case dengan gaya if else

// 	// switch {
// 	// case point == 8:
// 	// 	fmt.Println("Perfect")
// 	// case (point < 8) && (point > 3):
// 	// 	fmt.Println("Awesome")
// 	// 	fallthrough
// 	// case point < 5:
// 	// 	fmt.Println("You need to learn more")
// 	// default:
// 	// 	{
// 	// 		fmt.Println("Not Bad")
// 	// 		fmt.Println("You can be better!")
// 	// 	}
// 	// }

// 	// -- PERULANGAN -- (for, foreach, while)

// 	// var i = 0
// 	// for i < 5 {
// 	// 	fmt.Println("Angka ", i)
// 	// 	i++
// 	// }

// 	// perulangan tampa argumen

// 	// for {
// 	// 	fmt.Println("Angka ", i)
// 	// 	i++

// 	// 	if i == 6 {
// 	// 		break
// 	// 	}
// 	// }

// 	// penggunaan for - range

// 	// var xs = "123" // string
// 	// for i, v := range xs {
// 	// 	fmt.Println("Index=", i, "Value=", v)
// 	// }

// 	// var ys = [5]int{10, 20, 30, 40, 50} // array
// 	// for _, v := range ys {
// 	// 	fmt.Println("Value=", v)
// 	// }

// 	// var zs = ys[0:2] // slice
// 	// for _, v := range zs {
// 	// 	fmt.Println("Value=", v)
// 	// }

// 	// var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
// 	// for k, v := range kvs {
// 	// 	fmt.Println("Key=", k, "Value=", v)
// 	// }

// 	// // boleh juga baik k dan atau v nya diabaikan
// 	// for range kvs {
// 	// 	fmt.Println("Done")
// 	// }

// 	// // selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
// 	// for i := range 5 {
// 	// 	fmt.Print(i) // 01234
// 	// }

// 	// // Pengunaan keyword break dan continue
// 	// for i := 1; i <= 10; i++ {
// 	// 	if i%2 == 1 {
// 	// 		continue
// 	// 	}

// 	// 	if i > 8 {
// 	// 		break
// 	// 	}

// 	// 	fmt.Println("Angka", i)
// 	// }

// 	// Perulangan bersarang
// 	// for i := 0; i < 5; i++ {
// 	// 	for j := i; j < 5; j++ {
// 	// 		fmt.Print(j, " ")
// 	// 	}

// 	// 	fmt.Println()
// 	// }

// 	// -- ARRAY & SLICE & MAP --

// 	// var name [4]string

// 	// name[0] = "Trafalgar"
// 	// name[1] = "D"
// 	// name[2] = "Water"
// 	// name[3] = "Law"

// 	// fmt.Println(name[0], name[1], name[2], name[3])

// 	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

// 	// fmt.Println("Jumlah element \t\t", len(fruits))
// 	// fmt.Println("Isi semua element \t", fruits)

// 	// var numbers = [...]int{2, 3, 2, 4, 3}

// 	// fmt.Println("data array \t:", numbers)
// 	// fmt.Println("jumlah elemen \t:", len(numbers))

// 	// perulangan array menggunakan keyword for
// 	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

// 	// for i := 0; i < len(fruits); i++ {
// 	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
// 	// }

// 	// perulangan array menggunakan keyword for - range
// 	// for i, fruit := range fruits {
// 	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
// 	// }

// 	// perulangan array menggunakan keyword for - range dengan mengguanakn underscore (_)
// 	// for i, fruit := range fruits {
// 	// 	fmt.Printf("nama buah : %s\n", fruit)
// 	// } // kode ini akan error, solusinya gunaka seperti ini :

// 	// for _, fruit := range fruits {
// 	// 	fmt.Printf("nama buah : %s\n", fruit)
// 	// }

// 	//// SLICE
// 	// var fruits = []string{"apple", "grape", "banana", "melon"}
// 	// var newFruits = fruits[0:2]

// 	// var aFruits = fruits[0:3]
// 	// var bFruits = fruits[1:4]

// 	// var aaFruits = aFruits[1:2]
// 	// var baFruits = bFruits[0:1]

// 	// fmt.Println(fruits)   // [apple grape banana melon]
// 	// fmt.Println(aFruits)  // [apple grape banana]
// 	// fmt.Println(bFruits)  // [grape banana melon]
// 	// fmt.Println(aaFruits) // [grape]
// 	// fmt.Println(baFruits) // [grape]

// 	// baFruits[0] = "pinnaple"

// 	// fmt.Println(fruits)   // [apple pinnaple banana melon]
// 	// fmt.Println(aFruits)  // [apple pinnaple banana]
// 	// fmt.Println(bFruits)  // [pinnaple banana melon]
// 	// fmt.Println(aaFruits) // [pinnaple]
// 	// fmt.Println(baFruits) // [pinnaple]

// 	// fmt.Println(len(fruits)) // len: 4
// 	// fmt.Println(cap(fruits)) // cap: 4

// 	// fmt.Println(len(aFruits)) // len: 3
// 	// fmt.Println(cap(aFruits)) // cap: 4

// 	// fmt.Println(len(bFruits)) // len: 3
// 	// fmt.Println(cap(bFruits)) // cap: 3

// 	// -- MAP
// 	// var chiken = map[string]int{
// 	// 	"January":  50,
// 	// 	"Februari": 70,
// 	// 	"Maret":    80,
// 	// 	"April":    60,
// 	// }

// 	// // for key, val := range chiken {
// 	// // 	fmt.Println(key, " \t:", val)
// 	// // }

// 	// fmt.Println(len(chiken))
// 	// fmt.Println(chiken)

// 	// delete(chiken, "January")

// 	// fmt.Println(len(chiken))
// 	// fmt.Println(chiken)

// 	var names = []string{"John", "Wick"}
// 	printMessage("halo", names)
// }

// -- FUNCTION --

// func main() {
// 	var names = []string{"John", "Wick"}
// 	printMessage("halo", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

// func main() {
// 	var randomValue int

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = randomizer.Int()%(max-min+1) + min
// 	return value
// }

// func main() {
// 	divideNumber(10, 2)
// 	divideNumber(4, 0)
// 	divideNumber(8, -4)
// }

// func divideNumber(m, n int) {
// 	if n == 0 {
// 		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
// 		return
// 	}

// 	var res = m / n
// 	fmt.Printf("%d / %d = %d\n", m, n, res)
// }

// <-- Fungsi Multiple Return -->
// func calculate(d float64) (float64, float64) {
// 	// hitung luas
// 	var area = math.Pi * math.Pow(d/2, 2)
// 	// hitung keliling
// 	var circumference = math.Pi * d

// 	// kembalikan 2 nilai
// 	return area, circumference
// }

// func main() {
// 	var diameter float64 = 15
// 	var luas, keliling = calculate(diameter)

// 	fmt.Printf("luas lingkaran\t\t: %.2f \n", luas)
// 	fmt.Printf("keliling lingkaran\t: %.2f \n", keliling)
// }

// operation boolean

// func main() {
// 	var nilaiAkhir = 90
// 	var absensi = 80

// 	var lulusUjian bool = nilaiAkhir >= 90
// 	var lulusAbsensi bool = absensi >= 80

// 	var lulus bool = lulusUjian && lulusAbsensi

// 	fmt.Println("Mahasiswa ini dinyatakan : ", lulus)
// }
