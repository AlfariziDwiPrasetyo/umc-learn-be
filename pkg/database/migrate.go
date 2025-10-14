package database

import (
	"log"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
)

func Migrate() {
	if err := DB.AutoMigrate(&users.User{}); err != nil {
		log.Fatal("Failed to migrate: ", err)
	}
}
