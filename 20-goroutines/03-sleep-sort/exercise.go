package sleepSort

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"time"
)

func reverseSleepSort(input []uint) []uint {
	n := len(input)
	if n == 0 {
		return []uint{}
	}

	out := make(chan uint, n)

	for _, v := range input {
		val := v 
		go func() {
			delay := time.Duration(500-10*val) * time.Millisecond
			time.Sleep(delay)
			out <- val
		}()
	}

	res := make([]uint, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, <-out)
	}
	return res
}