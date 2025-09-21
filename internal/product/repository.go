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
	query := `
		INSERT INTO products (name, description, price, stock, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`

	_, err := r.db.ExecContext(ctx, query,
		item.Name,
		item.Description,
		item.Price,
		item.Stock,
	)

	if err != nil {
		return err
	}

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
