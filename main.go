package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Printf("a*a+2*a*b+b*b = %d\n", a*a+2*a*b+b*b)
}
