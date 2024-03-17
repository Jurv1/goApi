package service

import "rest/models"

type UserService interface {
	Save(user models.User) models.User
}

type userService struct {
	users []models.User
}

func New() UserService {
	return &userService{}
}

func (service *userService) Save(user models.User) models.User {
	service.users = append(service.users, user)
	return user
}
