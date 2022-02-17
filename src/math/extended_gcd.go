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

type ExtendedGCD interface {
	Run(a, b int32)
	GetGCD() int32
	GetQuotients() []int32
	GetCoefficients() []int32
	Example()
	ViewSummary()
}

type worker struct {
	prev_r   int32
	prev_s   int32
	prev_t   int32
	r        int32
	s        int32
	t        int32
	quotient int32
	results  []int32
}

func NewWorker() *worker {
	return &worker{
		prev_s: 1,
		s:      0,
		prev_t: 0,
		t:      1,
	}
}

func (w worker) GetGCD() int32 {
	return w.prev_r
}

func (w worker) GetQuotients() []int32 {
	return []int32{w.t, w.s}
}

func (w worker) GetCoefficients() []int32 {
	return []int32{w.prev_s, w.prev_t}
}

func (w worker) ViewSummary() {
	fmt.Printf("Bezout coefficients: %d, %d\n", w.prev_s, w.prev_t)
	fmt.Printf("gcd: %d\n", w.prev_r)
	fmt.Printf("quotients by gcd: %d, %d\n", w.t, w.s)
}

func Run(a, b int32) {
	w := &worker{
		prev_r: a,
		r:      b,
		prev_s: 1,
		s:      0,
		prev_t: 0,
		t:      1,
	}

	for w.r != 0 {
		w.quotient = (w.prev_r / w.r)
		w.prev_r, w.r = w.r, w.prev_r-(w.quotient*w.r)
		w.prev_s, w.s = w.s, w.prev_s-(w.quotient*w.s)
		w.prev_t, w.t = w.t, w.prev_t-(w.quotient*w.t)
	}
	w.ViewSummary()
}

func Example() {
	var A, B int32 = 240, 46
	var egcd ExtendedGCD
	Run(A, B)
	egcd.ViewSummary()
}

func main() {
	Example()
}
