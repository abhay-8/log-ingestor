package database

import (
	"fmt"

	"github.com/abhay-8/log-ingestor/backend/models"
)

func AutoMigrateDatabase() {
	fmt.Println("Starting migration")

	DB.AutoMigrate(
		&models.User{},
	)

	fmt.Println("Migrating database finished")
}
