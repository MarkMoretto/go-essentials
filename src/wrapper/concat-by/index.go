package concatby

// Join two strings by provided delimiter.
// Delimiter can be set early and the final
// functionality used at some point later on.

import "fmt"

// General string concat function with delimiter.
func concatenateTwo(strA, strB, delim string) string {
    return strA + delim + strB
}

// Main closure (wrapper/decorator) for two strings.
func ConcatBy(delimiter string) func(string, string) string {
    return func(s1, s2 string) string {
        // Return our previously-defined function.
        return concatenateTwo(s1, s2, delimiter)
        
        // If method is simple, can always skip and implement
        // directly:
        // return s1 + delimiter + s2
    }
}

// Testing example.
func ExampleConcatBy() {
    cbh := ConcatByHyphen("-")
    fmt.Println(cbh("followed", "by"))
    fmt.Println(cbh("followed", ""))
    fmt.Println(cbh("", "by"))
    fmt.Println(cbh("", ""))
    // Output:
    // followed-by
    // followed-
    // -by
    // -
}
