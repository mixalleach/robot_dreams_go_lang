package main

import (
	"testing"
)

func TestFibonacciIterative(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{n: 0, expected: 0},
		{n: 1, expected: 1},
		{n: 2, expected: 1},
		{n: 3, expected: 2},
		{n: 4, expected: 3},
		{n: 5, expected: 5},
		{n: 6, expected: 8},
		{n: 7, expected: 13},
		{n: 10, expected: 55},
		{n: 15, expected: 610},
	}

	for _, tt := range tests {
		result := FibonacciIterative(tt.n)
		if result != tt.expected {
			t.Errorf("FibonacciIterative(%d) = %d; expected %d", tt.n, result, tt.expected)
		}
	}
}

func TestFibonacciRecursive(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{n: 0, expected: 0},
		{n: 1, expected: 1},
		{n: 2, expected: 1},
		{n: 3, expected: 2},
		{n: 4, expected: 3},
		{n: 5, expected: 5},
		{n: 6, expected: 8},
		{n: 7, expected: 13},
		{n: 10, expected: 55},
		{n: 15, expected: 610},
	}

	for _, tt := range tests {
		result := FibonacciRecursive(tt.n)
		if result != tt.expected {
			t.Errorf("FibonacciRecursive(%d) = %d; expected %d", tt.n, result, tt.expected)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{n: 1, expected: true},
		{n: 2, expected: true},
		{n: 3, expected: true},
		{n: 4, expected: false},
		{n: 5, expected: true},
		{n: 6, expected: false},
		{n: 7, expected: true},
		{n: 8, expected: false},
		{n: 9, expected: false},
		{n: 10, expected: false},
		{n: 11, expected: true},
	}

	for _, tt := range tests {
		result := IsPrime(tt.n)
		if result != tt.expected {
			t.Errorf("IsPrime(%d) = %v; expected %v", tt.n, result, tt.expected)
		}
	}
}

func TestIsBinaryPalindrome(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{n: 5, expected: true},
		{n: 7, expected: true},
		{n: 6, expected: false},
		{n: 10, expected: false},
	}

	for _, tt := range tests {
		result := IsBinaryPalindrome(tt.n)
		if result != tt.expected {
			t.Errorf("IsBinaryPalindrome(%d) = %v; expected %v", tt.n, result, tt.expected)
		}
	}
}

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{s: "[{}]", expected: true},
		{s: "[{()}]", expected: true},
		{s: "[{sdfsdf(#$53)sdf}]", expected: true},
		{s: "[{]}", expected: false},
		{s: "[{()}]", expected: true},
	}

	for _, tt := range tests {
		result := ValidParentheses(tt.s)
		if result != tt.expected {
			t.Errorf("ValidParentheses(%q) = %v; expected %v", tt.s, result, tt.expected)
		}
	}
}

func TestIncrement(t *testing.T) {
	tests := []struct {
		binString string
		expected  int
	}{
		{binString: "00000001", expected: 2},
		{binString: "00000010", expected: 3},
		{binString: "00001010", expected: 11},
		{binString: "00011100", expected: 29},
	}

	for _, tt := range tests {
		result := Increment(tt.binString)
		if result != tt.expected {
			t.Errorf("Increment(%q) = %d; expected %d", tt.binString, result, tt.expected)
		}
	}
}
