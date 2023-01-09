package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	// function akan tetap dijalankan walaupun function error
	defer logging()

	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}
