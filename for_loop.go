package main

import "fmt"

func main() {
	// for biasa
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// for iterasi array manual
	var months = [...]string{
		"Januari",
		"Feburari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	for i := 0; i < 12; i++ {
		fmt.Println(months[i])
	}

	// for range, iterasi array otomatis
	for index, month := range months {
		fmt.Println("index ke", index, "adalah", month)
	}
}
