package gcd_lcm

// Switch positions for two integer values.
func switcheroo(high, low *int) {
	if *high > *low {
		t := low
		low = high
		high = t
	}
}

// Greatest common divisor core function.
func gcd(a, b int) int {
	var tmp int

	// Ensure that a <= b
	switcheroo(&a, &b)

	for b != 0 {
		tmp = b
		b = a % b
		a = tmp
	}

	return a
}

/// Leat common multiple for two or more integers.
func LCM(x, y int, rest ...int) int {
	res := x * y / gcd(x, y)

	for i := range rest {
		res = LCM(res, rest[i])
	}

	return res
}

/// Recursive greatest common divisor for two or more integers.
func GCD(x, y int, other ...int) (result int) {
	result = gcd(x, y)

	for i := range other {
		result = GCD(result, other[i])
	}

	return
}
