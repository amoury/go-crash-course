package main

import "fmt"

func main() {
	// ids := []int{33, 22, 556, 23, 3534}

	// for _, id := range ids {
	// 	fmt.Printf("ID - %d, is of type %T \n", id, id)
	// }

	// Ranging over Map

	emails := map[string]string{
		"Ansar": "a@xyz.com",
		"Memon": "b@xyz.com",
	}

	for name, email := range emails {
		fmt.Printf("Here is %s's email - %s \n", name, email)
	}
}
