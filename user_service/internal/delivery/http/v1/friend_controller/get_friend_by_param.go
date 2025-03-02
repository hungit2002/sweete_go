package friend_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"sweete/user_service/internal/dto"
)

func (c *Controller) GetFriendByParam(ctx *gin.Context) {
	message := "get friend by param"

	var input dto.GetFriendByParamInputDTO
	if err := ctx.ShouldBindQuery(&input); err != nil {
		panic(err)
	}
	if err := validator.New().Struct(input); err != nil {
		panic(err)
	}

	friend, err := c.friendUseCase.GetFriendByParam(ctx, input)
	if err != nil {
		res := dto.BadRequestErrorResponse(err.Error(), err)
		panic(res)
	}
	res := dto.SuccessResponse(message, friend)
	ctx.JSON(http.StatusOK, res)
}
