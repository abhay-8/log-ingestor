package config

import (
	"github.com/abhay-8/log-ingestor/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     database.CONFIG.FRONTEND_URL,
		AllowHeaders:     "Origin, Accept, Authorisation",
		AllowMethods:     "GET",
		AllowCredentials: true,
	})
}
