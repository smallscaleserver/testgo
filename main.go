package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Print("Enter first number: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter second number: ")
	fmt.Scanf("%d", &b)
	fmt.Print("Enter third number: ")
	fmt.Scanf("%d", &c)

	result := a + b + c
	fmt.Printf("Result: %d + %d + %d = %d\n", a, b, c, result)
}
