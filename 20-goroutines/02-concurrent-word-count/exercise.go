package wordcount

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"strings"
	"sync"
)

// CountWords takes a slice of strings and returns a map of word -> count.
// It uses goroutines and channels to count words concurrently.
func CountWords(lines []string) map[string]int {
	type pair struct {
		word  string
		count int
	}

	out := make(chan pair)
	var wg sync.WaitGroup

	// launch a goroutine per line
	for _, line := range lines {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			words := strings.Fields(s)
			local := make(map[string]int)
			for _, w := range words {
				local[w]++
			}
			for w, c := range local {
				out <- pair{w, c}
			}
		}(line)
	}

	// closer
	go func() {
		wg.Wait()
		close(out)
	}()

	// aggregate into final map
	result := make(map[string]int)
	for p := range out {
		result[p.word] += p.count
	}

	return result
}