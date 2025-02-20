package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"sweete/user_service/internal/dto"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Header("Content-Type", "application/json")
			if httpError, ok := err.(*dto.HttpError); ok {
				ctx.AbortWithStatusJSON(httpError.Meta.Code, httpError)
				return
			}

			if errors.Is(err.(error), gorm.ErrRecordNotFound) {
				res := dto.NoContentHttpResponse(nil, err.(error))
				ctx.AbortWithStatusJSON(res.Meta.Code, res)
				return
			}

			if validationError, ok := err.(validator.ValidationErrors); ok {
				res := dto.BadRequestErrorResponse(nil, validationError)
				ctx.AbortWithStatusJSON(res.Meta.Code, res)
				return
			}

			httpError := dto.InternalErrorResponse(nil, err.(error))
			ctx.AbortWithStatusJSON(httpError.Meta.Code, httpError)
			return
		}
	}()

	ctx.Next()
}
