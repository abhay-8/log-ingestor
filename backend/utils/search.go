package utils

import (
	"time"

	"github.com/abhay-8/log-ingestor/backend/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func logSearch(db *gorm.DB, field, value string) *gorm.DB {
	if value != "" {
		return db.Where(field+" LIKE ?", "%"+value+"%")
	}
	return db
}

func timestampSearch(db *gorm.DB, start, end string) *gorm.DB {
	if start != "" && end != "" {
		startTime, err := time.Parse(config.RFC3339, start)
		if err != nil {
			config.Logger.Warnw("Could not parse start time", "Error", err)
			return db
		}

		endTime, err := time.Parse(config.RFC3339, end)
		if err != nil {
			config.Logger.Warnw("Could not parse end time", "Error", err)
			return db
		}

		return db.Where("timestamp BETWEEN ? AND ?", startTime, endTime)
	} else if start != "" {
		startTime, err := time.Parse(config.RFC3339, start)
		if err != nil {
			config.Logger.Warnw("Could not parse start time", "Error", err)
			return db
		}

		return db.Where("timestamp >= ?", startTime)
	} else if end != "" {
		endTime, err := time.Parse(config.RFC3339, end)
		if err != nil {
			config.Logger.Warnw("Could not parse end time", "Error", err)
			return db
		}

		return db.Where("timestamp <= ?", endTime)
	}
	return db
}

func Search(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		fields := []string{
			"message",
			"level",
			"resource_id",
			"trace_id",
			"span_id",
			"commit",
			"parent_resource_id",
		}

		for _, field := range fields {
			value := c.Query(field, "")
			db = logSearch(db, field, value)
		}

		startTime := c.Query("start", "")
		endTime := c.Query("end", "")
		db = timestampSearch(db, startTime, endTime)

		return db
	}
}
