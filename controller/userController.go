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

func (controller *controller) UpdateUser(id string, ctx *gin.Context) {
	var user models.User

	ctx.BindJSON(&user)
	controller.service.Update(id, user)

}
