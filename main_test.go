package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		a, b, c  int
		expected int
	}{
		{2, 3, 4, 9},
		{0, 5, 1, 6},
		{-2, 4, 3, 5},
		{10, 10, 10, 30},
	}

	for _, tt := range tests {
		result := calculate(tt.a, tt.b, tt.c)
		if result != tt.expected {
			t.Errorf("calculate(%d, %d, %d) = %d; want %d", tt.a, tt.b, tt.c, result, tt.expected)
		}
	}
}
