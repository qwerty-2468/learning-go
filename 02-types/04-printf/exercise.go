package printer

import (
	"fmt"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

func printBool(b bool) string {
	return fmt.Sprintf("type: boolean value: %t", b)
}

func printInt(i int) string {
	return fmt.Sprintf("type: integer value: %d", i)
}

func printHex(i int) string {
	return fmt.Sprintf("type: integer in hexadecimal form value: %x", i)
}

func printFloat(f float64) string {
	return fmt.Sprintf("type: float value: %.2f", f)
}

func printString(s string) string {
	return fmt.Sprintf("type: string value: %q", s)
}

func concatStrings(a, b string) string {
	return a + b
}

func printConcatStrings(a, b string) string {
	return printString(concatStrings(a, b))
}
