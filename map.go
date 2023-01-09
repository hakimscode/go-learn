package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Heri",
		"job":  "Software Engineer",
	}
	person["gender"] = "Male"

	fmt.Println(person["name"])
	fmt.Println(person["job"])
	fmt.Println(person["gender"])

	// function len
	fmt.Println(len(person))

	// function make
	book := make(map[string]string)
	book["title"] = "Laskar Pelangi"
	book["author"] = "Andrea Hirata"
	book["year"] = "2005"
	fmt.Println(book)
	fmt.Println(book["title"])
	fmt.Println(book["author"])

	// function delete
	delete(book, "year")
	fmt.Println(book)
}
