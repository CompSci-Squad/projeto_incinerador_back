package main

import (
	"goapi/internal/middleware"
	"goapi/internal/routes"
	"goapi/pkg/configs"
	"goapi/pkg/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title Go Api Documentation
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

	// Capture system signals to properly shutdown consumers
	stop := make(chan os.Signal, 1)
	defer signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
