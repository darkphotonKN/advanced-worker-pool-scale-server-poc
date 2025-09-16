package product

import "github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/workerpool"

type Job struct {
	workerpool.Job
	service Service
}

func NewJob(service Service) workerpool.JobProcessor {
	return &Job{
		service: service,
	}
}
