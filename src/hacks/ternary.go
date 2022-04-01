package hacks

// Demo of a ternary operator in Go using a map object.

import "fmt"

func TernaryDemo() {
	var a, b int = 2, 4
    
	ternaryMax := map[bool]int{
		true:  a,
		false: b,
	}[a >= b]
    
	ternaryMin := map[bool]int{
		true:  a,
		false: b,
	}[a < b]
    
	fmt.Printf("Max: %d\nMin: %d\n", ternaryMax, ternaryMin)
    
}
