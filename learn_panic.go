package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		// aplikasi akan berhenti jika terjadi error
		panic("Aplikasi ERROR!!!")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
