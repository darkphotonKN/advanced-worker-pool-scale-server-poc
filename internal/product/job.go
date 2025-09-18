package product

import (
	"context"
	"time"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/model"
	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/workerpool"
	"github.com/google/uuid"
)

type ProductJob struct {
	workerpool.Job
	Data    *Product
	service Service
}

func NewJob(service Service) workerpool.JobProcessor {
	return &ProductJob{
		Job: workerpool.Job{
			ID:        uuid.New(),
			ResultCh:  make(chan model.Result),
			CreatedAt: time.Now(),
			Context:   context.Background(),
		},
		service: service,
	}
}

func (j *ProductJob) Execute() (interface{}, error) {
	switch j.Name {
	case "create":
		j.service.Create(j.Context, j.Data)
	}

	return nil, nil
}
