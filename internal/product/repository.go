package product

import (
	"context"
	"github.com/jmoiron/sqlx"
)

// repository struct - concrete implementation
type repository struct {
	db *sqlx.DB
}

// NewRepository returns the concrete type, not an interface
func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, item *Product) error {
	// Implementation left blank as requested
	return nil
}

func (r *repository) GetByID(ctx context.Context, id int) (*Product, error) {
	// Implementation left blank as requested
	return nil, nil
}

func (r *repository) List(ctx context.Context) ([]Product, error) {
	// Implementation left blank as requested
	return nil, nil
}

func (r *repository) Update(ctx context.Context, item *Product) error {
	// Implementation left blank as requested
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	// Implementation left blank as requested
	return nil
}

