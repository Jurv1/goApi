package controller

import (
	"github.com/gin-gonic/gin"
	"rest/cmd/service"
	"rest/models"
)

type UserController interface {
	CreateNewUser(ctx *gin.Context) models.User
}

type controller struct {
	service service.UserService
}

func New(userService service.UserService) UserController {
	return &controller{service: userService}
}

func (controller *controller) CreateNewUser(ctx *gin.Context) models.User {
	var user models.User

	ctx.BindJSON(&user)
	controller.service.Save(user)

	return user
}
