package product

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Create(ctx context.Context, item *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	List(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id int, item *Product) error
	Delete(ctx context.Context, id int) error
}

type Handler struct {
	service Service
}

// NewHandler accepts the Service interface
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	// Implementation left blank as requested
}

func (h *Handler) Get(c *gin.Context) {
	// Implementation left blank as requested
}

func (h *Handler) List(c *gin.Context) {
	// Implementation left blank as requested
}

func (h *Handler) Update(c *gin.Context) {
	// Implementation left blank as requested
}

func (h *Handler) Delete(c *gin.Context) {
	// Implementation left blank as requested
}
