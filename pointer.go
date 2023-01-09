package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// pointer in function
func ChangeAddress(address *Address) {
	address.Country = "France"
}

// pointer in method
func (address *Address) LocalProduct() {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Medan"

	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	ChangeAddress(address2)
	// change the original variable, cause pointing to the same memory address
	fmt.Println(address1)
	fmt.Println(address2)

	product1 := Address{"Banda Aceh", "Aceh", "Singapore"}
	fmt.Println(product1)
	product1.LocalProduct()
	fmt.Println(product1)
	fmt.Println(address2)

}
