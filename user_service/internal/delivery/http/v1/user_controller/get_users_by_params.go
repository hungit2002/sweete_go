package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"sweete/user_service/internal/dto"
)

func (c *Controller) GetUsersByParams(ctx *gin.Context) {
	message := "get users by params"

	var input dto.UserParamsDTO
	if err := ctx.ShouldBindQuery(&input); err != nil {
		panic(err)
	}
	if err := validator.New().Struct(input); err != nil {
		panic(err)
	}

	users, err := c.userUseCase.GetUsersByParams(ctx, input)
	if err != nil {
		res := dto.BadRequestErrorResponse(err.Error(), err)
		panic(res)
	}
	res := dto.SuccessResponse(message, users)
	ctx.JSON(200, res)
}
