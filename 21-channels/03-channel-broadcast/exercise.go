package channelbroadcaster

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import "context"

// channelBroadcast broadcasts all values from input to every channel in outputs.
// It stops and closes all output channels when the context is cancelled or
// the input channel is closed.
func channelBroadcast(ctx context.Context, input <-chan any, outputs []chan<- any) {
	go func() {
		// Ensure all outputs get closed when we are done.
		defer func() {
			for _, ch := range outputs {
				close(ch)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-input:
				if !ok {
					// input closed: stop broadcasting
					return
				}

				// Send v to all output channels.
				for _, ch := range outputs {
					select {
					case <-ctx.Done():
						return
					case ch <- v:
						// sent; ok even if receiver not ready immediately
					}
				}
			}
		}
	}()
}