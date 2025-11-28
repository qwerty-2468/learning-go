package functional

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func filter[E any](values []E, pred func(E) bool) []E {
	var out []E
	for _, v := range values {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}