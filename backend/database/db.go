package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// ConnectDB initializes the database connection pool
func ConnectDB() *pgxpool.Pool {
	// Load .env file
	err := godotenv.Load("../.env") // Adjust path based on where you run the app
	if err != nil {
		// Fallback if running from root
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("Warning: .env file not found, relying on system environment variables")
		}
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	// Configure connection pool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("❌ Unable to parse DATABASE_URL: %v\n", err)
	}

	// Pool settings (adjust for production later)
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	// Create connection pool
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("❌ Unable to connect to database: %v\n", err)
	}

	// Ping to verify connection
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("❌ Database ping failed: %v\n", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL (Neon)")
	return pool
}