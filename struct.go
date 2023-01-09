package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

// struct method
func (customer Customer) payBills(amount int) {
	fmt.Println(customer.Name, "pay the bill with amount", amount)
}

func main() {
	customer1 := Customer{
		Name:    "Heri",
		Address: "Pancoran",
		Age:     28,
	}

	customers := []Customer{
		{
			Name:    "Heri",
			Address: "Pancoran",
			Age:     29,
		},
	}

	fmt.Println(customer1)
	fmt.Println(customers[0].Name)

	// call struct method
	customer1.payBills(5000000)
}
