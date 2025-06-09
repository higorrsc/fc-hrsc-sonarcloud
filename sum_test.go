package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{2, 2, 4},    // Test sum of two positive numbers
		{-2, -2, -4}, // Test sum of two negative numbers
		{2, -2, 0},   // Test sum of a positive and a negative number
		{0, 0, 0},    // Test sum of two zeros
		{10, 5, 15},  // Additional test case
		{-10, 5, -5}, // Additional test case
	}

	for _, tt := range tests {
		if got := sum(tt.a, tt.b); got != tt.want {
			t.Errorf("sum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
