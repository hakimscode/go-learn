package main

import "fmt"

func sayHello(name string) string {
	return "Hello " + name
}

func addAndMinus(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func getCompleteName() (firstName, lastName string) {
	firstName = "Heri"
	lastName = "Hakim"
	return
}

// variadic function
func sumAll(calculate string, numbers ...int) int {
	total := 0

	switch calculate {
	case "sum":
		for _, number := range numbers {
			total += number
		}
	case "minus":
		for _, number := range numbers {
			total -= number
		}
	}

	return total
}

// type declaration in function
type Filter func(string) string

// function as a parameter
func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func filterWords(word string) string {
	switch word {
	case "anjing":
		return "..."
	case "babi":
		return "..."
	default:
		return word
	}
}

func main() {
	hello := sayHello("Heri Hakim")
	fmt.Println(hello)

	add, minus := addAndMinus(10, 7)
	fmt.Println(add, minus)

	a, b := getCompleteName()
	fmt.Println(a, b)

	total := sumAll("sum", 10, 5, 3, 4, 2)
	total2 := sumAll("minus", 10, 5, 3, 4, 2)
	fmt.Println(total, total2)

	// variadic function with slice
	arrNum := []int{10, 20, 30, 40, 50}
	total3 := sumAll("sum", arrNum...)
	total4 := sumAll("minus", arrNum...)
	fmt.Println(total3, total4)

	// function as a value
	sum := sumAll
	fmt.Println(sum("sum", arrNum...))

	// function as a parameter
	name := "Heri"
	sayHelloWithFilter(name, filterWords)
}
