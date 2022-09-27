package stringoperations

import "testing"

// Testing example.
func ExampleConcatByHyphen() {
    ConcatByHyphen := ConcatBy("-")
    fmt.Println(ConcatByHyphen("followed", "by"))
    fmt.Println(ConcatByHyphen("followed", ""))
    fmt.Println(ConcatByHyphen("", "by"))
    fmt.Println(ConcatByHyphen("", ""))
    // Output:
    // followed-by
    // followed-
    // -by
    // -
}
