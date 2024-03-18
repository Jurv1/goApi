package service

import (
	"rest/models"
	"rest/repository"
)

type QuestService interface {
	SaveQuest(quest models.Quest) models.Quest
	UpdateQuest(id string, quest models.Quest)
}

type questService struct {
	quests    []models.Quest
	questRepo repository.QuestRepository
}

func NewQuestService(repo repository.QuestRepository) QuestService {
	return &questService{
		questRepo: repo,
	}
}

func (questService questService) SaveQuest(quest models.Quest) models.Quest {
	questService.questRepo.SaveQuest(quest)
	return quest
}

func (questService questService) UpdateQuest(id string, user models.Quest) {
	questService.questRepo.UpdateQuest(id, user)
}
