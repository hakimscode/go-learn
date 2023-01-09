package main

import "fmt"

func main() {
	name := "Hakim"

	switch name {
	case "Heri":
		fmt.Println(name + " adalah admin")
	case "Hakim":
		fmt.Println(name + " adalah manager")
	default:
		fmt.Println("Failed")
	}

	// short statement di switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama kepanjangan")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa ekspresi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama kepanjangan")
	case length < 2:
		fmt.Println("Nama kependekan")
	default:
		fmt.Println("Nama sudah benar")
	}
}
