package main

import (
	"github.com/abhay-8/log-ingestor/backend/config"
	"github.com/abhay-8/log-ingestor/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	database.LoadEnv()
	database.ConnectToDB()
	database.ConnectToCache()
	database.AutoMigrate()

	config.AddLogger()
}

func main() {
	defer config.LoggerCleanUp()
	app := fiber.New(fiber.Config{
		ErrorHandler: fiber.DefaultErrorHandler,
		BodyLimit:    config.BODY_LIMIT,
	})

	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(config.CORS())
	// app.Use(config.RATE_LIMITER())

	app.Listen(":" + database.CONFIG.PORT)
}
