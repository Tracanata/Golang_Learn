package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("Fatra")

	if data == nil {
		fmt.Println("Data Map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
