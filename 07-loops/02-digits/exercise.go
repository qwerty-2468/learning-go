package digits

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func multiplyDigits(n int) int {
	prod := 1

	for n > 0 {
		d := n % 10   // last digit
		prod *= d     // multiply into product
		n /= 10       // drop last digit
	}

	return prod
}