package workerpool

import (
	"sync"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/model"
)

type Pool struct {
	jobs        chan Job // the channel of jobs that workers pull from
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
	maxWorkerCount := 20 // TODO: update to a dynamic count based on CPU cycles
	safeBufferSize := maxWorkerCount * bufferMultiplier

	newPool := Pool{
		jobs:        make(chan Job, safeBufferSize),
		noOfWorkers: maxWorkerCount,
		wg:          sync.WaitGroup{},
	}

	// start all the worker goroutines
	for i := 0; i < maxWorkerCount; i++ {
		go newPool.worker()
	}

	return &newPool
}

/**
* Primary definition of a worker, the method that carries out the work sent to the
* jobs channel.
**/
func (p *Pool) worker() {
	for job := range p.jobs {
		result, err := job.Execute()

		// parse incoming request and pass it to work on the correct service and method

		if err != nil {
			job.ResultCh <- model.Result{
				Result: nil,
				Error:  &err,
			}
		}

		job.ResultCh <- model.Result{
			Result: result,
			Error:  nil,
		}
	}
}
