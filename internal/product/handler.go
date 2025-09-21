package product

import (
	"context"
	"net/http"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/workerpool"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	pool    *workerpool.Pool
	service HandlerService
}

type HandlerService interface {
	Create(ctx context.Context, item *Product) error
	GetByID(ctx context.Context, id int) (*Product, error)
	List(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, id int, item *Product) error
	Delete(ctx context.Context, id int) error
}

func NewHandler(pool *workerpool.Pool, service HandlerService) *Handler {
	return &Handler{pool: pool, service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job := NewJob(h.service, "create")
	job.SetData(&product)

	// queue into job channel to queue work
	h.pool.Submit(job)

	// listen for result after business logic is done from service.go
	result := <-job.GetResultCh()

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "sucessfully created new product."})
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
