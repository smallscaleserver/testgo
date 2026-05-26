package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter first number (a): ")
	scanner.Scan()
	a, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	fmt.Print("Enter second number (b): ")
	scanner.Scan()
	b, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	fmt.Print("Enter third number (c): ")
	scanner.Scan()
	c, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)

	result := math.Pow(a, b) + c
	fmt.Printf("Result: %.0f^%.0f + %.0f = %.0f\n", a, b, c, result)
}
