package main

import "fmt"

func calculate(a, b, c int) int {
	return a + b + c
}

func main() {
	var a, b, c int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Print("Enter c: ")
	fmt.Scan(&c)
	fmt.Printf("a + b + c = %d\n", calculate(a, b, c))
}
