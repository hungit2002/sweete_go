package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"sweete/user_service/internal/dto"
)

func (c *Controller) Register(ctx *gin.Context) {
	message := "Register user_controller"

	var input dto.RegisterDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		panic(err)
	}
	if err := validator.New().Struct(input); err != nil {
		panic(err)
	}

	user, accessToken, err := c.authUseCase.Register(ctx, input.FirstName, input.Surname, input.DateOfBirth, input.Gender, input.Phone, input.Email, input.Password)
	if err != nil {
		res := dto.BadRequestErrorResponse(message, err)
		panic(res)
	}

	registerResponse := dto.AuthResponseDTO{
		AccessToken: accessToken,
		User:        *user,
	}
	res := dto.SuccessResponse(message, registerResponse)
	ctx.JSON(http.StatusOK, res)
}
