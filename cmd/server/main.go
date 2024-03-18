package main

import (
	"github.com/gin-gonic/gin"
	"rest/controller"
	"rest/repository"
	"rest/service"
)

var (
	userRepo       repository.UserRepository = repository.NewUserRepository()
	userService    service.UserService       = service.New(userRepo)
	userController controller.UserController = controller.New(userService)
)

func main() {
	app := gin.Default()

	app.POST("/user", func(context *gin.Context) {
		user := userController.CreateNewUser(context)

		context.JSON(200, gin.H{
			"resp": user,
		})
	})

	app.PUT("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		userController.UpdateUser(id, context)

		context.JSON(204, gin.H{
			"resp": "Ok",
		})
	})

	app.Run(":8080")
}
