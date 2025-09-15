package workerpool

import (
	"fmt"
	"sync"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/product"
)

type Pool struct {
	jobs    chan Job // the channel of jobs that workers pull from
	workers int
	wg      sync.WaitGroup // for synchronizing workers

	// DI of services the worker has to interface with
	productService product.ServiceInterface
}

const maxWorker = 200

/**
* initializes all workers and goroutines from the get-go.
* this keeps all workers available for the entire lifetime of
* the pool, and at least as part of the goal, the application.
**/
func NewPool() {
	// start all the worker goroutines

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
			job.ResultCh <- err
		}

		job.ResultCh <- result
	}

}

func (p *Pool) ListenForJobs() {
	go func() {
		for {
			select {
			case job := <-p.jobs:
				p.WorkOnJob(job)

			}
		}

	}()
}

func (p *Pool) AddWorker(job Job) error {
	return nil
}

func (p *Pool) WorkOnJob(job Job) error {
	// add a worker if there is still worker capacity available
	if maxWorker >= p.workers {
		return fmt.Errorf("Workers are all exhausted.")
	}

	return p.AddWorker(job)
}
