package sleepSort

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"time"
)

// reverseSleepSort returns the input uint-slice sorted in the reverse order.
func reverseSleepSort(input []uint) []uint {
	n := len(input)
	if n == 0 {
		return []uint{}
	}

	out := make(chan uint, n)

	for _, v := range input {
		val := v // copy loop variable
		go func() {
			// input is in [1,50]; reverse sleep: 500 - x*10 ms
			delay := time.Duration(500-10*val) * time.Millisecond
			time.Sleep(delay)
			out <- val
		}()
	}

	res := make([]uint, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, <-out)
	}
	// because larger numbers sleep less, they arrive first => already reverse-sorted
	return res
}