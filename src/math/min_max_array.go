package main

import "fmt"

// Integer slice type.
type IntArr []int

// Return length of slice.
func (ia IntArr) Len() int { return len(ia) }

// Return max value of slice.
func (ia IntArr) Max() int {
	var maxN int = -1
	for l, r := 0, ia.Len()-1; l < r; l, r = l+1, r-1 {
		if ia[l] > ia[r] {
			if ia[l] > maxN {
				maxN = ia[l]
			}
		} else if ia[r] > maxN {
			maxN = ia[r]
		}
	}
	return maxN
}

// Return min value of slice.
const (
	MaxInt32 = 1<<31 - 1
	MaxInt64 = 1<<63 - 1
)

func (ia IntArr) Min() int {
    // This is for 32-bit operations, but switching
    // 31 to 63
	var minN int = MaxInt32
	for l, r := 0, ia.Len()-1; l < r; l, r = l+1, r-1 {
		if ia[l] < ia[r] {
			if ia[l] < minN {
				minN = ia[l]
			}
		} else if ia[r] < minN {
			minN = ia[r]
		}
	}
	return minN
}

func MinMaxDemo() {
	var arr IntArr = IntArr{4, 10, 5, 4, 2, 10}
	fmt.Printf("Max value: %d\n", arr.Max())
	fmt.Printf("Min value: %d\n", arr.Min())
}
