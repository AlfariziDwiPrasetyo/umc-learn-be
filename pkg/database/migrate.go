package database

import (
	"log"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/posts"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&users.User{}, &authentications.Authentications{}, &posts.Post{}, &comments.Comments{}); err != nil {
		log.Fatal("Failed to migrate: ", err)
	}
}
