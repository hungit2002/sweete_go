package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"sweete/user_service/internal/dto"
)

func (c *Controller) GetUserDetail(ctx *gin.Context) {
	message := "get user detail"

	var input dto.UserIdDTO
	if err := ctx.ShouldBind(&input); err != nil {
		panic(err)
	}
	if err := validator.New().Struct(input); err != nil {
		panic(err)
	}

	user, err := c.userUseCase.GetUserDetail(ctx, input.ID)
	if err != nil {
		res := dto.BadRequestErrorResponse(message, err)
		panic(res)
	}
	res := dto.SuccessResponse(message, user)
	ctx.JSON(http.StatusOK, res)
}
