package channelmultiplexer

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"context"
	"sync"
)

// ChannelMultiplex multiplexes multiple input channels into a single output channel.
func ChannelMultiplex(ctx context.Context, inputs []chan any) chan any {
	out := make(chan any) // unbuffered; callers can wrap with buffering if needed

	var wg sync.WaitGroup
	wg.Add(len(inputs))

	// Start one goroutine per input channel.
	for _, ch := range inputs {
		in := ch // capture loop variable
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- v:
						// delivered
					}
				}
			}
		}()
	}

	// Close output when all readers are done or context is cancelled.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}