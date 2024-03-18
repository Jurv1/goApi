package service

import (
	"rest/models"
	"rest/repository"
)

type UserService interface {
	SaveUser(user models.User) models.User
	UpdateUser(id string, user models.User)
}

type userService struct {
	users    []models.User
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepo: repo,
	}
}

func (service userService) SaveUser(user models.User) models.User {
	service.userRepo.Save(user)
	return user
}

func (service userService) UpdateUser(id string, user models.User) {
	service.userRepo.Update(id, user)
}
