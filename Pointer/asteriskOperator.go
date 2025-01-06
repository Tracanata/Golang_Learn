package main

import "fmt"

type Alamat struct {
	Kota, Provinsi, Negara string
}

func main() {
	address1 := Alamat{"Cianjur", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.Kota = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	//address2 = &Alamat{"Jakarta", "DKI JAKARTA", "Indonesia"}
	*address2 = Alamat{"Jakarta", "DKI JAKARTA", "Indonesia"}
	//address1 & address2 akan mengikuti data alamat yang terkini
	fmt.Println(address1)
	fmt.Println(address2)

}
