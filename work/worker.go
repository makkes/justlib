// Package work contains routines to asynchronously dispatch jobs to
// a fixed set of Goroutines. Use this package to e.g. create a worker with 10
// Goroutines and distribute work among them.
package work

import "errors"

// Payload wraps the data that represents a job to be processed.
type Payload struct {
	Data interface{}
}

// Worker represents a worker unit that distributes jobs between its working
// Goroutines.
type Worker struct {
	workerCount int
	workerFunc  func(Payload)
	dispatcher  chan Payload  // channel for dispatching jobs to the worker
	completions chan struct{} // channel with one entry per completed job
	quit        chan struct{} // channel used to signal the workers to exit
	inProgress  int
	shutdown    bool
}

// NewWorker creates a new worker that maintains exactly workerCount
// Goroutines. Each Goroutine calls workerFunc for processing the given data.
func NewWorker(workerCount int, workerFunc func(Payload)) *Worker {
	dispatcher := make(chan Payload)
	completions := make(chan struct{})
	quit := make(chan struct{})

	// spawn worker goroutines
	for i := 0; i < workerCount; i++ {
		go func() {
			for {
				select {
				case job := <-dispatcher:
					workerFunc(job)
					select {
					case completions <- struct{}{}:
					default:
					}
				case <-quit:
					return
				}
			}
		}()
	}

	worker := &Worker{
		workerCount,
		workerFunc,
		dispatcher,
		completions,
		quit,
		0,
		false,
	}

	return worker

}

// Completions returns a channel that is sent a value to every time when a job
// is completed. This lets you keep track of completed jobs. Be aware that you
// need to start reading from this channel before dispatching jobs to the
// Worker, otherwise you would not receive all completions.
func (w *Worker) Completions() <-chan struct{} {
	return w.completions
}

// Dispatch feeds a new job to the Worker. If no idle Goroutine is available,
// this function blocks until the job can be processed.
func (w *Worker) Dispatch(job Payload) error {
	if w.shutdown {
		return errors.New("WORKER IS SHUT DOWN")
	}
	w.dispatcher <- job
	return nil
}

// Quit waits for all workers to complete their jobs and returns afterwards.
// After this function returns, all worker routines are stopped and you cannot
// use this worker, anymore.
func (w *Worker) Quit() {
	for i := 0; i < w.workerCount; i++ {
		w.quit <- struct{}{}
	}
	w.shutdown = true
}
