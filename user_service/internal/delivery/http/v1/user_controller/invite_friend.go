package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"sweete/user_service/internal/dto"
)

func (c *Controller) InviteFriend(ctx *gin.Context) {
	message := "invite friend"

	var input dto.InputInviteFriendDTO
	if err := ctx.ShouldBind(&input); err != nil {
		panic(err)
	}

	if err := validator.New().Struct(input); err != nil {
		panic(err)
	}

	err := c.userUseCase.InviteFriend(ctx, input.UserID, input.FriendID, input.Status)
	if err != nil {
		res := dto.BadRequestErrorResponse(err.Error(), err)
		panic(res)
	}
	res := dto.SuccessResponse(message, nil)
	ctx.JSON(200, res)
}
