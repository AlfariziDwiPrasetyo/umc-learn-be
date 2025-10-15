package authentications

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func (h *Handler) Refresh(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	var request struct {
// 		RefreshToken string `json:"refresh_token"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error":   true,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// }
