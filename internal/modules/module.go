package modules

import (
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	authHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/authentications"
	commentHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/comments"
	likeHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/likes"
	postHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/posts"
	userHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/users"
	authRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/authentications"
	commentRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/comments"
	likeRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/likes"
	postRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/posts"
	userRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/users"
	authService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/authentications"
	commentService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/comments"
	likeService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/likes"
	postService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/posts"
	userService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/users"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAll(r *gin.Engine, db *gorm.DB, cfg *configs.Config, cld *cloudinary.Cloudinary) {
	// User
	userRepo := userRepository.NewRepository(db)
	userSvc := userService.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(r, cfg, userSvc)

	userHandler.RegisterRoute()

	// Authentication
	authRepo := authRepository.NewRepository(db)
	authSvc := authService.NewService(cfg, userRepo, authRepo)
	authHandler := authHandler.NewHandler(r, authSvc)

	authHandler.RegisterRoute()

	// Post
	postRepo := postRepository.NewRepository(db)
	postSvc := postService.NewService(cfg, postRepo, cld)
	postHandler := postHandler.NewHandler(r, cfg, postSvc)

	postHandler.RegisterRoute()

	// Comment
	commentRepo := commentRepository.NewRepository(db)
	commentSvc := commentService.NewService(cfg, commentRepo, cld)
	commentHandler := commentHandler.NewHandler(r, cfg, commentSvc)

	commentHandler.RegisterRoute()

	// Likes
	likeRepo := likeRepository.NewRepository(db)
	likeSvc := likeService.NewService(cfg, likeRepo)
	likeHandler := likeHandler.NewHandler(r, likeSvc, cfg)

	likeHandler.RegisterRoute()

}
