package work_test

import "fmt"
import "strconv"
import "time"
import "github.com/justsocialapps/justlib/work"

// this is our worker function that is called for every new job
func f(p work.Payload) interface{} {
	fmt.Println(p.Data)
	time.Sleep(200 * time.Millisecond)
	return nil
}

func Example() {

	// create a worker with 3 goroutines that can process jobs in parallel
	worker := work.NewWorker(3, f, false)

	// create 100 jobs and dispatch them to the worker.
	for i := 0; i < 100; i++ {
		// this call will block when all 3 goroutines are currently busy.
		worker.Dispatch(work.Payload{Data: strconv.Itoa(i)})
	}

	// this call makes sure that the worker stops all goroutines as soon as
	// they have processed all remaining jobs.
	worker.Quit()

}
