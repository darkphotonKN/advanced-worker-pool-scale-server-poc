package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/darkphotonKN/advanced-worker-pool-scale-server-poc/config"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Seed products
	log.Println("Starting product seeding...")
	start := time.Now()

	if err := seedProducts(db); err != nil {
		log.Fatal("Failed to seed products:", err)
	}

	elapsed := time.Since(start)
	log.Printf("Successfully seeded 100,000 products in %v", elapsed)
}

func seedProducts(db *sqlx.DB) error {
	ctx := context.Background()

	// Configuration
	batchSize := 5000 // Optimal batch size for bulk inserts
	totalProducts := 100000

	// Start transaction
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Process in batches for memory efficiency
	for batch := 0; batch < totalProducts; batch += batchSize {
		// Calculate batch boundaries
		start := batch + 1
		end := batch + batchSize
		if end > totalProducts {
			end = totalProducts
		}

		// Build bulk insert query with multiple value sets
		var valueStrings []string
		var valueArgs []interface{}

		for i := start; i <= end; i++ {
			// Calculate position for placeholders
			pos := (i - start) * 6
			valueStrings = append(valueStrings,
				fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
					pos+1, pos+2, pos+3, pos+4, pos+5, pos+6))

			// Add values for this product
			valueArgs = append(valueArgs,
				fmt.Sprintf("TestProduct%d", i),               // name
				"This is a test product for seeding purposes", // description
				99.99,      // price
				100,        // stock
				time.Now(), // created_at
				time.Now(), // updated_at
			)
		}

		// Execute bulk insert
		query := fmt.Sprintf(`
			INSERT INTO products (name, description, price, stock, created_at, updated_at)
			VALUES %s`, strings.Join(valueStrings, ","))

		if _, err := tx.ExecContext(ctx, query, valueArgs...); err != nil {
			return fmt.Errorf("failed to insert batch starting at %d: %w", start, err)
		}

		// Log progress
		log.Printf("Inserted %d/%d products...", end, totalProducts)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

