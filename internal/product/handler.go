package product

import (
	"context"

	"github.com/gin-gonic/gin"
)

// ServiceInterface defined where it's consumed (in handler)
// Handler decides what methods it needs from the service
type ServiceInterface interface {
	Create(ctx context.Context, item *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	List(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id int, item *Product) error
	Delete(ctx context.Context, id int) error
}

type Handler struct {
	service ServiceInterface
}

// NewHandler accepts the interface it defined
func NewHandler(service ServiceInterface) *Handler {
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

