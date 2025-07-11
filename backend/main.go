// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hnifmaghfur/bookingtogo/config"
	"github.com/hnifmaghfur/bookingtogo/internal/repository"
	"github.com/hnifmaghfur/bookingtogo/internal/service"
	"github.com/hnifmaghfur/bookingtogo/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file for main app
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to database
	db := config.ConnectDB()

	// Initialize repositories
	customerRepo := repository.NewCustomerRepository(db)
	familyListRepo := repository.NewFamilyListRepository(db)
	nationalityRepo := repository.NewNationalityRepository(db)

	// Initialize services
	customerService := service.NewCustomerService(customerRepo)
	familyListService := service.NewFamilyListService(familyListRepo, customerRepo)
	nationalityService := service.NewNationalityService(nationalityRepo)

	// Initialize handlers
	customerHandler := routes.NewCustomerHandler(customerService)
	familyListHandler := routes.NewFamilyListHandler(familyListService)
	nationalityHandler := routes.NewNationalityHandler(nationalityService)

	// Setup router
	router := mux.NewRouter()
	routes.SetupRoutes(router, customerHandler, familyListHandler, nationalityHandler)

	// Start server
	port := ":8080"
	fmt.Printf("Server berjalan di port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
