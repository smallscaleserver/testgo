package main

import (
	"bufio"
	"fmt"
	"math/big"
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

	fmt.Print("Enter third number (c): ")
	scanner.Scan()
	cStr := strings.TrimSpace(scanner.Text())

	// Parse a as big.Int
	a := new(big.Int)
	a.SetString(aStr, 10)

	// Parse b as int (exponent must be integer)
	b, _ := strconv.Atoi(bStr)

	// Parse c as big.Int
	c := new(big.Int)
	c.SetString(cStr, 10)

	// Calculate a^b using big.Int.Exp()
	result := new(big.Int)
	result.Exp(a, big.NewInt(int64(b)), nil)

	// Add c to result
	result.Add(result, c)

	fmt.Printf("Result: %s^%d + %s = %s\n", aStr, b, cStr, result.String())
}
