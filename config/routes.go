package config

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/middleware"
	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/product"
	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/internal/workerpool"
)

func SetupRoutes(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	// initalize workerpool
	pool := workerpool.NewPool()

	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo)

	productHandler := product.NewHandler(pool, productService)

	// API routes
	api := router.Group("/api")
	{
		// Public routes
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "healthy"})
		})

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/products", productHandler.List)
			protected.POST("/products", productHandler.Create)
			protected.GET("/products/:id", productHandler.Get)
			protected.PUT("/products/:id", productHandler.Update)
			protected.DELETE("/products/:id", productHandler.Delete)
		}
	}

	return router
}

