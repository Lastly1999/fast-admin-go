package services

import (
	"go-service/models"
	"go-service/repository"
)

type IUserService interface {
	Get(user models.User) (*models.User, error)
}

type UserService struct {
	Repo repository.IUserRepository
}

func (sev *UserService) Get(user models.User) (*models.User, error) {
	return sev.Repo.Get(user)
}
