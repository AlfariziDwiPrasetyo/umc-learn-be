package authentications

import (
	"context"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
	"github.com/gin-gonic/gin"
)

type authService interface {
	SignUp(ctx context.Context, req authentications.RegisterUser) error
	SignIn(ctx context.Context, req authentications.LoginUser) (*authentications.Tokens, error)
	Refresh(ctx context.Context, req authentications.RefreshTokenRequest) (*authentications.Tokens, error)
	LogOut(ctx context.Context, refreshToken string) error
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
	r.POST("refresh", h.Refresh)
	r.POST("logout", h.LogOut)
}
