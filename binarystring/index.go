package binarystring

// Handles conversion of binary string to decimal value.

import (
    "fmt"
)

func DemoBinaryToDecimalX() {
	testCases := []struct {
		binStr   string
		expected int
	}{
		{`10`, 2},
		{`110`, 6},
		{`10011100010000`, 10_000},
	}
	for _, tc := range testCases {
		actual := BinaryToDecimalX(tc.binStr)
		fmt.Printf("Actual: %d vs %d expected\n", actual, tc.expected)
	}
}

// Convert strng of binary to integer value.
// Handles both odd and even length strings, so they
// don't have to be divisible by 2.
func BinaryToDecimalX(str string) (tot int) {
	var n, i, j, lhs, rhs int
	n = len(str)
	for i, j = n-1, 0; i > j; i, j = i-1, j+1 {
		lhs, rhs = binToDec(str[j], i), binToDec(str[i], j)
		tot = tot + lhs + rhs
	}
    // If odd number of elements, calculate the middle value and
    // increment total by result.
	if n%2 != 0 && i == j {
		tot = tot + binToDec(str[i], i)
	}
	return
}

// Convert byte to binary position integer value.
func binToDec(b byte, expon int) int {
	return byteToInt(b) * 1 << expon
}

// Convert decimal byte to int.
func byteToInt(b byte) int {
	return int(b - '0')
}
