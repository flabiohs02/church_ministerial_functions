package main

import (
	"log"

	"church_ministerial_functions/config"
	"church_ministerial_functions/domain"
	"church_ministerial_functions/handler"
	"church_ministerial_functions/repository"
	"church_ministerial_functions/routers"
	"church_ministerial_functions/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := config.InitDb()

	// Auto-migrate the MinisterialFunction schema
	db.AutoMigrate(&domain.MinisterialFunction{})

	// Dependency Injection
	mfRepo := repository.NewGormMinisterialFunctionRepository(db)
	mfService := usecase.NewMinisterialFunctionService(mfRepo)
	mfHandler := handler.NewMinisterialFunctionHandler(mfService)

	// Set up Fiber app
	app := fiber.New()

	// Setup routes
	routers.SetupMinisterialFunctionRoutes(app, mfHandler)

	log.Println("Server listening on :8686")
	if err := app.Listen(":8686"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
