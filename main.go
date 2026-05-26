package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Enter first number (a): ")
	fmt.Scanf("%f", &a)
	fmt.Print("Enter second number (b): ")
	fmt.Scanf("%f", &b)
	fmt.Print("Enter third number (c): ")
	fmt.Scanf("%f", &c)

	result := math.Pow(a, b) + c
	fmt.Printf("Result: %.0f^%.0f + %.0f = %.0f\n", a, b, c, result)
}
