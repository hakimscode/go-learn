package main

import "fmt"

func main() {

	// deklarasi array kosong manual
	var names [3]string
	names[0] = "Heri"
	names[1] = "Hakim"
	names[2] = "Setiawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi array langsung dengan value
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// function len, menghitung panjang dari array
	fmt.Println(len(names))
	fmt.Println(len(values))

}
