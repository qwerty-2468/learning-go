package channelmultiplexer

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"context"
	"sync"
)

func channelMultiplex(ctx context.Context, inputs []chan any) chan any {
	out := make(chan any)

	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, ch := range inputs {
		in := ch 
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
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}