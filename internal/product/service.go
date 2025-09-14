package product

import (
	"context"
)

// RepositoryInterface defined where it's consumed (in service)
type RepositoryInterface interface {
	Create(ctx context.Context, item *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	List(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, item *Product) error
	Delete(ctx context.Context, id int) error
}

// service struct - concrete implementation
type service struct {
	repo RepositoryInterface
}

// NewService returns the concrete type, not an interface
// This follows the pattern where service doesn't need to follow an interface
func NewService(repo RepositoryInterface) *service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, item *Product) error {
	// Implementation left blank as requested
	return nil
}

func (s *service) GetByID(ctx context.Context, id int) (*Product, error) {
	// Implementation left blank as requested
	return nil, nil
}

func (s *service) List(ctx context.Context) ([]Product, error) {
	// Implementation left blank as requested
	return nil, nil
}

func (s *service) Update(ctx context.Context, id int, item *Product) error {
	// Implementation left blank as requested
	return nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	// Implementation left blank as requested
	return nil
}

