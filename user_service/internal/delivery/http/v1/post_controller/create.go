package post_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"sweete/user_service/internal/dto"
)

func (c *Controller) Create(ctx *gin.Context) {
	message := "create post"

	var input dto.CreatePostInputDTO

	if err := ctx.ShouldBind(&input); err != nil {
		panic(err)
	}
	if err := validator.New().Struct(input); err != nil {
		panic(err)
	}

}
