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

// Euclidean GCD
func EuclideanGCD(a, b uint32) uint32 {
	if a%b == 0 {
		return b
	}
	return EuclideanGCD(b, a%b)
}

// Binary GCD
// Note: BinaryGCD(x, y, 1) ==> GCD(x, y)
func BinaryGCD(a, b, res uint32) uint32 {
	switch {
	case a == b:
		return res * a
	case a%2 == 0:
		switch b%2 == 0 {
		case true:
			return BinaryGCD(a/2, b/2, 2*res)
		default:
			return BinaryGCD(a/2, b, res)
		}
	case b%2 == 0:
		return BinaryGCD(a, b/2, res)
	case a > b:
		return BinaryGCD(a-b, b, res)
	default:
		return BinaryGCD(a, b-a, res)
	}
}
