package main

import (
	"crud-redis/app"
	"crud-redis/config"
	"crud-redis/repository"
	"crud-redis/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	// Initialize Redis client
	err = config.RedisInit(cfg.RedisAddress, cfg.RedisPassword)
	if err != nil {
		fmt.Printf("Failed to initialize Redis: %v\n", err)
		return
	}
	redis := config.RedisConnect()

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		fmt.Printf("Failed to initialize database: %v\n", err)
		return
	}
	defer db.Close()

	// Initialize repository and service
	repo := repository.NewRepository(redis, db)
	service := app.SetupApp(db, repo)
	e := echo.New()
	e.Use(middleware.Logger())

	routes.ApiRoutes(e, service)

	e.Use(middleware.Recover())

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
