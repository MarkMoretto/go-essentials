package join

// Premade join function variables.
var (
    JoinByHyphen = JoinBy("-")   // Hyphen.
    JoinByComma = JoinBy(",")    // Comma.
    JoinByCommaWS = JoinBy(", ") // Comma and whitespace.
)

// Join two or more strings by a specified delimiter.
func JoinBy(delim string) func(...string) string {
	return func(strs ...string) string {
		outs := strs[0]
		for _, s := range strs[1:] {
			outs += delim + s
		}
		return outs
	}
}
