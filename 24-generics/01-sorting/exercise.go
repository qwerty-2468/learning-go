package sorting

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"sort"
)

type Ordered interface {
	~int64 | ~float64 | ~string
}

func sortSlice[T Ordered](input []T) []T {
	out := make([]T, len(input))
	copy(out, input)

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})
	return out
}
