package main

/**
Some utilities for fast IO when doing competitive coding/code challenges.
*/

import (
	"bufio"
	"fmt"
	"os"
)

// Write to stdout.
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func scanf(f string, a ...interface{}) {
	fmt.Fscanf(reader, f, a...)
}

// Read from Stdin
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
func printf(f string, a ...interface{}) {
	fmt.Fprintf(writer, f, a...)
}


func main() {
	//! Need to manually clear write buffer.
	defer writer.Flush()
	
	var s string // String
	
	// var x, y int // Numeric Int
	// var i, j float32 // Numeric Float
	
	//! Use scanf with relevant signature (eg - %d for decimal, %f for float)
	//! Pass pointer
	scanf("%s\n", &s)
	printf("%s\n", s)
	
	//! Integer example
	// scanf("%d %d\n", &x, &y)
	// printf("%d %d\n", y, x)
}
