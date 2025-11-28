package search

import (
	"io"
	"bufio"
	"regexp"
	"strings"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// search reads a text and a word and returns true if the word appears in the text and false if it does not.
func contain(reader io.Reader, word string) bool {
	// INSERT YOUR CODE HERE
	// read all text
	buf := bufio.NewScanner(reader)
	buf.Split(bufio.ScanLines)

	var sb strings.Builder
	for buf.Scan() {
		sb.WriteString(buf.Text())
		sb.WriteByte(' ')
	}
	text := sb.String()

	// remove punctuation, keep letters/digits and spaces
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	text = re.ReplaceAllString(text, " ")

	text = strings.ToLower(text)
	target := strings.ToLower(word)

	words := strings.Fields(text)

	count := 0
	for _, w := range words {
		if w == target {
			count++
			if count > 1 {
				return true
			}
		}
	}
	return false
}
