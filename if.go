package main

import "fmt"

func main() {
	name := "Heri"

	if name == "Heri" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("Not Authorized")
	}

	// short statement di if
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama available")
	}
}
