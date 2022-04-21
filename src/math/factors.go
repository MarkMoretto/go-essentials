package math

import "fmt"

func DemoFactors() {
    // Factors of N
	var N int = 20
    
    // Prepped integer array
	var arr = makeArray(20)
    
    // Factor array result
	var facts = factors(N, arr)
    
    // Print results to stdout.
    fmt.Printf("%v\n", facts)
}

// Main (prime) factors function
func factors(n int, intArray []int) []int {
	var outf []int
	for intArray[n] > 0 {
		outf = append(outf, intArray[n])
		n /= intArray[n]
	}
	outf = append(outf, n)
	return outf
}

// Prep factor helper array
func makeArray(n int) []int {
	var iarr = make([]int, n+1)
	var i int = 2
	var k int
	for (i * i) <= n {
		if iarr[i] == 0 {
			k = i * i
			for k <= n {
				if iarr[k] == 0 {
					iarr[k] = i
				}
				k += i
			}
		}
		i++
	}
	return iarr
}
