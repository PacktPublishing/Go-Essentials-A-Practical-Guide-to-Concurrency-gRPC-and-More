package main

import (
	"testing"
)

// func TestSum(t *testing.T){
// 	result := Sum(3,5)
// 	expect := 8

// 	if result != expect {
// 		t.Errorf("Sum(3,5) return %d, expected %d", result, expect)
// 	}
// }

func TestSum(t *testing.T){
	testCases := []struct {
		a, b, expected int
	}{
		{3, 5, 8},
		{-2, 6, 4},
		{0, 0, 0},
	}

	for _, tc := range testCases {
		result := Sum(tc.a, tc.b)

		if result != tc.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}