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

func ConnectDB() *pgxpool.Pool {
	// .env file root folder mein hai, isliye path change kiya gaya hai
	err := godotenv.Load("../../.env") 
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			err = godotenv.Load(".env")
		}
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("❌ Unable to parse DATABASE_URL: %v\n", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("❌ Unable to connect to database: %v\n", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("❌ Database ping failed: %v\n", err)
	}

	fmt.Println("✅ Successfully connected to PostgreSQL (Neon)")
	return pool
}