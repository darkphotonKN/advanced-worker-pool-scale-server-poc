package workerpool

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	id        uuid.UUID
	Data      interface{} // holds all the original data for the work
	ResultCh  chan interface{}
	CreatedAt time.Time
}

/**
* What workers should run, tells that what to do.
**/
func (j *Job) Execute(data interface{}) (interface{}, error) {
	return nil, nil
}

/**
* Tells worker where to send the result of executing the job.
**/
func (j *Job) HandleResult(data interface{}, err error) (interface{}, error) {
	return nil, nil
}
