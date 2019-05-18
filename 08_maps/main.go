package main

import "fmt"

func main() {
	// Define map

	x := "Bob"
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Bill"] = "bill@gmail.com"
	emails["Rob"] = "rob@gmail.com"

	fmt.Println(emails[x])

	delete(emails, "Bob")

	cities := map[string]string{
		"Bob":  "Blore",
		"Rob":  "Japan",
		"Bill": "Dubai",
	}

	fmt.Println(cities)
}
