package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message : ", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi ERROR!!!")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	// aplikasi akan tetap berjalan walaupun panic karena telah menggunakan recover
	fmt.Println("Done")
}
