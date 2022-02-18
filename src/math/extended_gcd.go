package math

// NOT FULLY IMPLEMENTED.

/*
Extended Greatest Common Divisor
- or -
Extended Euclidean Algorithm

*/

import (
	"fmt"
)

// Worker struct for attributes that will be used in the algorithm.
// An instance of this will be returned when the main EGCD function is executed.
type worker struct {
	prev_r int32
	prev_s int32
	prev_t int32
	r int32
	s int32
	t int32
	quotient int32
}

// Retrieve GCD.
func (w worker) GetGCD() int32 {
	return w.prev_r
}

// Retrieve quotients
func (w worker) GetQuotients() []int32 {
	return []int32{w.t, w.s}
}

// Retrieve coefficients
func (w worker) GetCoefficients() []int32 {
	return []int32{w.prev_s, w.prev_t}
}

// View summary of results.
func (w worker) ViewSummary() {
	fmt.Printf("Bézout coefficients: %d, %d\n", w.prev_s, w.prev_t)
	fmt.Printf("gcd: %d\n", w.prev_r)
	fmt.Printf("quotients by gcd: %d, %d\n", w.t, w.s)
}

// Driver function. This returns the worker instance, which can
// be used to retrieve various results of the algorithm, as noted above.
func EGCD(a, b int32) *worker {
	w := &worker{
		prev_r: a,
		r: b,
		prev_s: 1,
		s: 0,
		prev_t: 0,
		t: 1,		
	}

	for w.r != 0 {
		w.quotient = (w.prev_r / w.r)
		w.prev_r, w.r = w.r, w.prev_r-(w.quotient*w.r)
		w.prev_s, w.s = w.s, w.prev_s-(w.quotient*w.s)
		w.prev_t, w.t = w.t, w.prev_t-(w.quotient*w.t)
	}
	// w.ViewSummary()
	return w
}

func Example() {
	var A, B int32 = 240, 46
	res := EGCD(A, B)
	res.ViewSummary()
}

func main() {
	Example()
}
