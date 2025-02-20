package auth_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sweete/user_service/internal/dto"
)

func (c *Controller) Login(ctx *gin.Context) {
	message := "Login successfully"

	var input dto.LoginDTO
	if err := ctx.ShouldBind(&input); err != nil {
		log.Error(err)
		panic(err)
	}
	if err := validator.New().Struct(input); err != nil {
		log.Error(err)
		panic(err)
	}
	user, token, err := c.authUseCase.Login(ctx, input.Email, input.Phone, input.Password)
	if err != nil {
		res := dto.BadRequestErrorResponse(err.Error(), err)
		panic(res)
	}
	loginResponse := dto.AuthResponseDTO{
		AccessToken: token,
		User:        *user,
	}
	res := dto.SuccessResponse(message, loginResponse)
	ctx.JSON(http.StatusOK, res)
}
