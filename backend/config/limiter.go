package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RATE_LIMITER() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        2000,
		Expiration: 1 * time.Hour,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests, try again later.",
			})
		},
	})
}

const LIMIT = 5 * 1024 * 1024 // 5 MB
