package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sweete/core/pkg/jwt"
	"sweete/user_service/config"
	"sweete/user_service/internal/dto"
)

func CheckPermission(ctx *gin.Context) {
	str := ctx.GetHeader("Authorization")
	if len(str) == 0 {
		str = ctx.GetHeader("token")
	}
	if len(str) == 0 {
		str = ctx.PostForm("token")
	}

	if len(str) > 0 && strings.HasPrefix(str, "Bearer ") {
		str = strings.TrimPrefix(str, "Bearer ")
	}

	if len(str) == 0 {
		log.Error("unauthorized")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse(http.StatusUnauthorized, "unauthorized", nil))
		return
	}

	if str == config.GetInstance().App.TokenToServer {
		ctx.Next()
		return
	}
	jwtService := jwt.NewJWTService()
	user, err := jwtService.VerifyToken(str)
	if err != nil {
		log.Error("unauthorized")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse(http.StatusUnauthorized, "unauthorized", nil))
		return
	}
	ctx.Set("user_id", user.ID)
	ctx.Set("user_controller", user)

	ctx.Next()
}
