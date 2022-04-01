package hacks

// Demo of a ternary operator in Go using a map object.

import "fmt"

type Ternary map[bool]int

func TernaryDemo() {
	var a, b int = 2, 4
    
	ternaryMax := Ternary{
		true:  a,
		false: b,
	}[a >= b]
    
	ternaryMin := Ternary{
		true:  a,
		false: b,
	}[a < b]
    
	fmt.Printf("Max: %d\nMin: %d\n", ternaryMax, ternaryMin)
    
}
