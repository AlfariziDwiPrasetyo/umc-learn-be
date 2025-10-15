package authentications

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"github.com/gin-gonic/gin"
)

type authService interface {
	SignUp(ctx context.Context, req users.RegisterUser) error
	SignIn(ctx context.Context, req users.LoginUser) (*authentications.Tokens, error)
}

type Handler struct {
	*gin.Engine
	AuthSvc authService
}

func NewHandler(api *gin.Engine, authSvc authService) *Handler {
	return &Handler{
		Engine:  api,
		AuthSvc: authSvc,
	}
}

func (h *Handler) RegisterRoute() {
	r := h.Group("auth")

	r.POST("sign-up", h.SignUp)
	r.POST("sign-in", h.SignIn)
}
