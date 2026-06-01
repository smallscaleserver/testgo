package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter first number (a): ")
	scanner.Scan()
	aStr := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter second number (b): ")
	scanner.Scan()
	bStr := strings.TrimSpace(scanner.Text())

	// Parse a as int
	a, _ := strconv.Atoi(aStr)

	// Parse b as int
	b, _ := strconv.Atoi(bStr)

	// Calculate a + b
	result := a + b

	fmt.Printf("%s + %s = %d\n", aStr, bStr, result)
}
