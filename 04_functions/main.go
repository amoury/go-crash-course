package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Ansar"))
	fmt.Println(getSum(20, 10))
}
