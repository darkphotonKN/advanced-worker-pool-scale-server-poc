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
	service JobService
}

type JobService interface {
	Create(ctx context.Context, item *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	List(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id int, item *Product) error
	Delete(ctx context.Context, id int) error
}

func NewJob(service JobService, name string) workerpool.JobProcessor {
	return &ProductJob{
		Job: workerpool.Job{
			ID:        uuid.New(),
			ResultCh:  make(chan model.Result),
			Name:      name,
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
