package threadpool

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
import (
	"context"
	"log"
	"sync"
)

type Runnable interface {
	Run(context.Context) error
}

type ThreadPool interface {
	Run(Runnable)
	Close()
}

type threadPool struct {
	ctx       context.Context
	cancel    context.CancelFunc
	tasks     chan Runnable
	errCh     chan error
	wg        sync.WaitGroup
	closeOnce sync.Once
}

func NewThreadPool(n int) (ThreadPool, chan error) {
	ctx, cancel := context.WithCancel(context.Background())

	p := &threadPool{
		ctx:    ctx,
		cancel: cancel,
		// buffered so Run() is non-blocking but still bounded
		tasks: make(chan Runnable, n),
		// buffered to avoid blocking on error reporting
		errCh: make(chan error, n*2),
	}

	for i := 0; i < n; i++ {
		p.wg.Add(1)
		go p.worker()
	}

	return p, p.errCh
}

func (p *threadPool) worker() {
	defer p.wg.Done()

	for {
		select {
		case <-p.ctx.Done():
			return
		case r, ok := <-p.tasks:
			if !ok {
				return
			}
			if r == nil {
				continue
			}
			// RUN TASK HERE, inside worker goroutine
			if err := r.Run(p.ctx); err != nil {
				select {
				case p.errCh <- err:
				default:
					log.Println("threadpool error:", err)
				}
			}
		}
	}
}

// Run submits a task; it must NOT start a new goroutine.
func (p *threadPool) Run(r Runnable) {
	select {
	case <-p.ctx.Done():
		return
	default:
		select {
		case p.tasks <- r:
		case <-p.ctx.Done():
		}
	}
}

func (p *threadPool) Close() {
	p.closeOnce.Do(func() {
		p.cancel()
		close(p.tasks)
		p.wg.Wait()
		close(p.errCh)
	})
}