package scanning

import (
	"io"
	"bufio"
	"unicode"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// counter reads a text and returns the counted values.
func counter(reader io.Reader) int {
	// INSERT YOUR CODE HERE
	r := bufio.NewReader(reader)
	count := 0

	for {
		ch, _, err := r.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		if unicode.IsLower(ch) {
			count++
		}
	}
	return count
}
