package database

import (
	"log"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&users.User{}); err != nil {
		log.Fatal("Failed to migrate: ", err)
	}
}
