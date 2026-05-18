package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Printf("a + b * a = %d\n", a+b*a)
}
