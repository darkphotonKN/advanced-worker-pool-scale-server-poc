package workerpool

import (
	"context"
	"fmt"
	"time"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/model"
	"github.com/google/uuid"
)

/**
* Represents a baseline Job's functionality.
* the interface abstraction of what it means to be a job, to be able to execute the task it carries
**/
type JobProcessor interface {
	Execute() (interface{}, error)
	GetResultCh() chan model.Result
}

/*
* Baseline job type that would be embedded by all other concrete jobs
**/
type Job struct {
	ID       uuid.UUID
	ResultCh chan model.Result

	// name of the operation (corresponding method) to run from that domain's service
	Name string

	// the payload for carrying out the job
	Data      interface{}
	CreatedAt time.Time
	Context   context.Context
}

/**
* What workers should run, tells that worker what to do.
**/
func (j *Job) Execute() (interface{}, error) {
	fmt.Printf("\nJob Executed\n\n")
	return nil, nil
}

/**
* Tells worker where to send the result of executing the job.
**/
func (j *Job) HandleResult(data interface{}, err error) (interface{}, error) {
	return nil, nil
}

/**
* Provides cleaner access to the ResultCh since we are using an interface abstraction.
**/
func (j *Job) GetResultCh() chan model.Result {
	return j.ResultCh
}
