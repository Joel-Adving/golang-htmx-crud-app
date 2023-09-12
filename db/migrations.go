package db

import (
	"github.com/Joel-Adving/testing-stuff/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}
