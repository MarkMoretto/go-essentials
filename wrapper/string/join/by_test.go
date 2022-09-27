package join

import "testing"

// Example of JoinBy function using comma + horizontal whitespace delimiter.
func ExampleJoinByCommaWhiteSpace() {
	tokens := []string{"Hello", "beautiful", "world"}
	jbc := JoinBy(", ")
	result := jbc(tokens...)
	fmt.Println(result)
	// Output:
	// Hello, beautiful, world
}

// Example of JoinBy function using comma delimiter.
func ExampleJoinByComma() {
	tokens := []string{"Hello", "beautiful", "world"}
	jbc := JoinBy(",")
	result := jbc(tokens...)
	fmt.Println(result)
	// Output:
	// Hello,beautiful,world
}

// Example of JoinBy function using vertical bar.
func ExampleJoinByVBar() {
	tokens := []string{"Hello", "beautiful", "world"}
	jbc := JoinBy("|")
	result := jbc(tokens...)
	fmt.Println(result)
	// Output:
	// Hello|beautiful|world
}
