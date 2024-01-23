package routers

import (
	"github.com/abhay-8/log-ingestor/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func LogRouter(app *fiber.App) {
	logRoutes := app.Group("/logs")
	logRoutes.Post("/", controller.AddLog)
	logRoutes.Get("/search", controller.GetSearchLogs)
	logRoutes.Get("/", controller.GetAllLogs)
}
