package main

import (
	"clicknext-backend/docs"
	"clicknext-backend/internal/config"
	"clicknext-backend/internal/delivery/http/handlers"
	"clicknext-backend/internal/delivery/routes"
	"clicknext-backend/internal/infrastructure/database"
	"fmt"
	"log"

	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	db, err := database.NewGromDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting raw DB connection: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Database connection established successfully!")

	app := fiber.New()

	docs.SwaggerInfo.Title = "ClickNext Backend API"
	docs.SwaggerInfo.Description = "This is the API documentation for ClickNext backend service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	h := handlers.NewHandler(db)

	routes.RegisterRoutes(app, h)

	app.Listen(":" + cfg.Server.Address)
}
