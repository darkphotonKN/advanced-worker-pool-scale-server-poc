package workerpool

import (
	"sync"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/product"
)

type Pool struct {
	jobs        chan Job // the channel of jobs that workers pull from
	noOfWorkers int
	wg          sync.WaitGroup // for synchronizing workers

	// DI of services the worker has to interface with
	productService product.Service
}

const (
	maxWorkerCount   int = 20
	bufferMultiplier int = 2
)

/**
* initializes all workers and goroutines from the get-go.
* this keeps all workers available for the entire lifetime of
* the pool, and at least as part of the goal, the application.
**/
func NewPool() *Pool {
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
		result, err := job.Execute("placeholder")

		// parse incoming request and pass it to work on the correct service and method

		if err != nil {
			job.ResultCh <- err
		}

		job.ResultCh <- result
	}
}
