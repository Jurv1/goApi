package controller

import (
	"github.com/gin-gonic/gin"
	"rest/models"
	"rest/service"
)

type QuestController interface {
	CreateNewQuest(ctx *gin.Context) models.Quest
	UpdateQuest(id string, ctx *gin.Context)
}

type questController struct {
	service service.QuestService
}

func NewQuestController(questService service.QuestService) QuestController {
	return &questController{service: questService}
}

func (controller *questController) CreateNewQuest(ctx *gin.Context) models.Quest {
	var quest models.Quest

	ctx.BindJSON(&quest)
	controller.service.SaveQuest(quest)

	return quest
}

func (controller *questController) UpdateQuest(id string, ctx *gin.Context) {
	var quest models.Quest

	ctx.BindJSON(&quest)
	controller.service.UpdateQuest(id, quest)

}
