package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func fibonacci(n int) *big.Int {
	if n <= 0 {
		return big.NewInt(0)
	}
	if n == 1 {
		return big.NewInt(1)
	}

	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 2; i <= n; i++ {
		c := new(big.Int).Add(a, b)
		a.Set(b)
		b.Set(c)
	}

	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter first number (a): ")
	scanner.Scan()
	aStr := scanner.Text()

	fmt.Print("Enter second number (b): ")
	scanner.Scan()
	bStr := scanner.Text()

	// Parse a as big.Int with base 10
	a := new(big.Int)
	a.SetString(aStr, 10)

	// Parse b as big.Int with base 10
	b := new(big.Int)
	b.SetString(bStr, 10)

	// Calculate a + b
	result := new(big.Int).Add(a, b)

	fmt.Printf("%s + %s = %s\n", aStr, bStr, result.String())

	// Calculate fibonacci of the result
	n := int(result.Int64())
	if n < 0 {
		fmt.Println("Cannot calculate Fibonacci for negative numbers")
	} else if n > 10000 {
		fmt.Printf("Result too large for practical Fibonacci calculation (n=%d)\n", n)
	} else {
		fib := fibonacci(n)
		fmt.Printf("Fibonacci(%d) = %s\n", n, fib.String())
	}
}
