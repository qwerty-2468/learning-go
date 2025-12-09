package concurrentprimes

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"sort"
	"sync"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimes(n int) []int {
	if n < 2 {
		return []int{}
	}

	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 2; i <= n; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			if isPrime(x) {
				ch <- x
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var primes []int
	for p := range ch {
		primes = append(primes, p)
	}

	sort.Ints(primes)

	return primes
}