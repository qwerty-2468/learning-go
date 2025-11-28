package strings

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

// returns the literal multiline string:
// some
// multiline
// string
func multilineString() string {
	return `some
multiline
string`
}

// returns the length of the string
func stringLen(s string) int {
	return len(s)
}

// returns the input string without its first character
func trimFirstChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[1:]
}

// returns the input string without its last character
func trimLastChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[:len(s)-1]
}

// replaces the first character with 'A'
func swapFirstChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	return "A" + s[1:]
}

// replaces the last character with 'A'
func swapLastChar(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return "A"
	}
	return s[:len(s)-1] + "A"
}

// prepends 'A' to the string
func prependChar(s string) string {
	return "A" + s
}

// appends 'A' to the string
func appendChar(s string) string {
	return s + "A"
}
