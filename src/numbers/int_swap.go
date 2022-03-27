package numbers

// Simple function to take references of two integers and
// swap values if the first argument is greater
// than the second argument.
func IntSwap(first *int, second *int) {
	if *first > *second {
        // Dereference second value to pointer for later.
		pSecond := *second
		*second = *first
        // Set first memory address to value of second integer.
		*first = pSecond
	}
}

// IntSwap demo.
func DemoIntSwap() {
    // f and s are integers
    // For clarity:
    //      f => 'first'
    //      s => 'second'
	var f, s int
	f = 4
	s = 3
	fmt.Println("Before: ", f, s)
	IntSwap(&f, &s)
	fmt.Println("After: ", f, s)
}
