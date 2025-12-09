package subtask

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"context"
	"time"
)

func StartTask(ctx context.Context) (result string, err error) {
	subCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resCh := make(chan string)
	errCh := make(chan error)

	go func() {
		res, e := SubTask(subCtx)
		if e != nil {
			errCh <- e
			return
		}
		resCh <- res
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case e := <-errCh:
		return "", e
	case r := <-resCh:
		return "Main task status: " + r, nil
	}
}

func SubTask(ctx context.Context) (result string, err error) {
	resCh := make(chan string)
	errCh := make(chan error)

	go func() {
		timer := time.NewTimer(200 * time.Millisecond)
		defer timer.Stop()

		select {
		case <-ctx.Done():
			errCh <- ctx.Err()
		case <-timer.C:
			resCh <- "Subtask completed successfully"
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case e := <-errCh:
		return "", e
	case r := <-resCh:
		return r, nil
	}
}