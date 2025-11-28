package closures

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

import "errors"
func proxy(fn func(string) int) func(string) (int, error) {
	on := true

	return func(s string) (int, error) {
		if on {
			on = false
			return fn(s), nil
		}
		on = true
		return 0, errors.New("proxy is off")
	}
}