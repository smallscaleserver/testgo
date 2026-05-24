package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Printf("a * b = %d\n", multiply(a, b))
}
