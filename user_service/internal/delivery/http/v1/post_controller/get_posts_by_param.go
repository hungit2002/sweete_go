package post_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"sweete/user_service/internal/dto"
)

func (c *Controller) GetPostsByParam(ctx *gin.Context) {
	message := "get posts by param"

	var input dto.GetPostsByParamInputDTO
	if err := ctx.ShouldBindQuery(&input); err != nil {
		panic(err)
	}
	if err := validator.New().Struct(input); err != nil {
		panic(err)
	}

	posts, err := c.postUseCase.GetPostsByParam(ctx, input)
	if err != nil {
		res := dto.BadRequestErrorResponse(err.Error(), err)
		panic(res)
	}
	res := dto.SuccessResponse(message, posts)
	ctx.JSON(http.StatusOK, res)
}
