package controller

import (
	"github.com/gin-gonic/gin"
	"rest/models"
	"rest/service"
)

type UserController interface {
	CreateNewUser(ctx *gin.Context) models.User
	UpdateUser(id string, ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func New(userService service.UserService) UserController {
	return &userController{service: userService}
}

func (controller *userController) CreateNewUser(ctx *gin.Context) models.User {
	var user models.User

	ctx.BindJSON(&user)
	controller.service.SaveUser(user)

	return user
}

func (controller *userController) UpdateUser(id string, ctx *gin.Context) {
	var user models.User

	ctx.BindJSON(&user)
	controller.service.UpdateUser(id, user)

}
