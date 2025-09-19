package workerpool

import (
	"fmt"
	"sync"
	"time"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/model"
)

type Pool struct {
	jobs        chan JobProcessor // the channel of jobs that workers pull from
	noOfWorkers int
	wg          sync.WaitGroup // for synchronizing workers
}

const (
	bufferMultiplier int = 2
)

/**
* initializes all workers and goroutines from the get-go.
* this keeps all workers available for the entire lifetime of
* the pool, and at least as part of the goal, the application.
**/
func NewPool() *Pool {
	// TODO: update to a dynamic count based on CPU cycles
	maxWorkerCount := 20
	safeBufferSize := maxWorkerCount * bufferMultiplier

	newPool := Pool{
		jobs:        make(chan JobProcessor, safeBufferSize),
		noOfWorkers: maxWorkerCount,
		wg:          sync.WaitGroup{},
	}

	// start all the worker goroutines
	for i := 0; i < maxWorkerCount; i++ {
		go newPool.worker(i)
	}

	return &newPool
}

/**
* Primary definition of a worker, the method that carries out the work sent to the
* jobs channel.
**/
func (p *Pool) worker(workerNo int) {
	fmt.Println("Starting worker: ", workerNo)
	for job := range p.jobs {
		result, err := job.Execute()

		// parse incoming request and pass it to work on the correct service and method
		jobResultCh := job.GetResultCh()

		if err != nil {
			jobResultCh <- model.Result{
				Result: nil,
				Error:  &err,
			}
		}

		jobResultCh <- model.Result{
			Result: result,
			Error:  nil,
		}
	}
}

/**
* Allows for the queuing of a job onto the job channel.
**/
func (p *Pool) Submit(job JobProcessor) error {

	// NOTE: using select pattern here to validate message went through, not to listen to multiple
	// incoming channels
	select {
	case p.jobs <- job:
		return nil
	case <-time.After(time.Second * 1):
		return fmt.Errorf("timed out when sending job %+v to channel\n", job)
	}
}
