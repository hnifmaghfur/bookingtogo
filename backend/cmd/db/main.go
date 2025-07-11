// cmd/db/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hnifmaghfur/bookingtogo/config"
	"github.com/hnifmaghfur/bookingtogo/database/seeders"

	"github.com/joho/godotenv" // You'll need this here too
)

func main() {
	// Load .env file first, as config.GetMigrationURL might need env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to database
	db := config.ConnectDB()

	// Get migration URL from config
	migrationURL := config.GetMigrationURL()

	// Parse command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/db/main.go <command>")
		fmt.Println("Commands:")
		fmt.Println("  migrate: Runs all pending database migrations.")
		fmt.Println("  seed:    Runs database seeders (e.g., initial nationalities).")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "migrate":
		fmt.Println("Running database migrations...")
		m, err := migrate.New("file://database/migrations", migrationURL)
		if err != nil {
			log.Fatalf("Failed to initialize database migration: %v", err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run database migrations: %v", err)
		}
		if err == migrate.ErrNoChange {
			fmt.Println("No new migrations to apply.")
		} else {
			fmt.Println("Database migrations applied successfully!")
		}

	case "seed":
		fmt.Println("Running database seeders...")
		seeders.RunSQLSeeder(db, "database/seeders/initial_nationality_data.sql")
		fmt.Println("Initial nationalities seeded successfully.")

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: migrate, seed")
		os.Exit(1)
	}
}
