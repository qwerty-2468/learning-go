package pipeline

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func generator(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func multiplier(in <-chan int) <-chan float32 {
	out := make(chan float32)
	go func() {
		for n := range in {
			out <- float32(n * 5)
		}
		close(out)
	}()
	return out
}

func collector(in <-chan float32) []float32 {
	var result []float32
	for v := range in {
		result = append(result, v)
	}
	return result
}