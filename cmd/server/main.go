package main

import (
	"github.com/gin-gonic/gin"
	"rest/controller"
	"rest/repository"
	"rest/service"
)

var (
	userRepo        repository.UserRepository  = repository.NewUserRepository()
	userService     service.UserService        = service.NewUserService(userRepo)
	userController  controller.UserController  = controller.New(userService)
	questRepo       repository.QuestRepository = repository.NewQuestRepository()
	questService    service.QuestService       = service.NewQuestService(questRepo)
	questController controller.QuestController = controller.NewQuestController(questService)
)

func main() {
	app := gin.Default()

	app.POST("/user", func(context *gin.Context) {
		user := userController.CreateNewUser(context)

		context.JSON(201, gin.H{
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

	app.POST("/quest", func(context *gin.Context) {
		quest := questController.CreateNewQuest(context)

		context.JSON(201, gin.H{
			"resp": quest,
		})
	})

	app.Run(":8080")
}
