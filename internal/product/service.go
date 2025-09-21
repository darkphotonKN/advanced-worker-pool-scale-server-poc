package product

import (
	"context"
	"fmt"
)

type Repository interface {
	Create(ctx context.Context, item *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	List(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, item *Product) error
	Delete(ctx context.Context, id int) error
}

// service struct - concrete implementation
type service struct {
	repo Repository
}

// NewService returns the concrete type, not an interface
func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, item *Product) error {
	fmt.Printf("\nRecieved payload from worker: %+v\n\n", item)

	// Validate the product
	if err := item.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Call repository to persist to database
	if err := s.repo.Create(ctx, item); err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func (s *service) GetByID(ctx context.Context, id int) (*Product, error) {
	return nil, nil
}

func (s *service) List(ctx context.Context) ([]Product, error) {
	return nil, nil
}

func (s *service) Update(ctx context.Context, id int, item *Product) error {
	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return nil
}
