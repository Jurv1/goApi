package service

import (
	"rest/models"
	"rest/repository"
)

type UserService interface {
	Save(user models.User) models.User
	Update(id string, user models.User)
}

type userService struct {
	users    []models.User
	userRepo repository.UserRepository
}

func New(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (service userService) Save(user models.User) models.User {
	service.userRepo.Save(user)
	return user
}

func (service userService) Update(id string, user models.User) {
	service.userRepo.Update(id, user)
}
