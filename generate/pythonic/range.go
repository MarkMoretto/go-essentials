package pythonic

// Simple script to generate sequences of numbers.
// Possible usage:
//     seq := GenerateRange(10) // [0 1 2 3 4 5 6 7 8 9]
//     seq := GenerateRange(1, 11) // [1 2 3 4 5 6 7 8 9 10]
//     seq := GenerateRange(1, 11, 2) // [1 3 5 7 9]

// Note: Cannot take named arguments, yet.  So, if you want [0, 2, 4, 6, 8, 10]
// the propet call is: 
//     seq = GenerateRange(0, 11, 2)

import "fmt"

type Range struct {
	start, stop, stepSize int
}

func NewRange() *Range {
	return &Range{
		stepSize: 1,
	}
}

func (r Range) Size() int {
	return (r.stop - r.start) / r.stepSize
}

func GenerateRange(startStopStep ...int) []int {
	var seq = NewRange()

	switch len(startStopStep) {
	case 1:
		(*seq).stop = startStopStep[0]
	case 2:
		(*seq).start = startStopStep[0]
		(*seq).stop = startStopStep[1]
	case 3:
		(*seq).start = startStopStep[0]
		(*seq).stop = startStopStep[1]
		(*seq).stepSize = startStopStep[2]
	default:
		return nil
	}

	var outs = make([]int, 0, seq.Size())

	for i := seq.start; i < seq.stop; i += seq.stepSize {
		outs = append(outs, i)
	}
	return outs
}

func RangeDemo() {
	r1 := GenerateRange(10)
	fmt.Println(r1)
	r2 := GenerateRange(1, 11)
	fmt.Println(r2)
	r3 := GenerateRange(1, 10, 2)
	fmt.Println(r3)    
}
