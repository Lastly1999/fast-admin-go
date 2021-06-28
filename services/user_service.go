package services

import (
	"errors"
	"go-service/models"
	"go-service/repository"
)

type IUserService interface {
	Get(user models.User) (*models.User, *models.UserRole, error)
}

type UserService struct {
	UserRepo repository.IUserRepository
	RoleRepo repository.IRoleRepository
}

func (sev *UserService) Get(user models.User) (*models.User, *models.UserRole, error) {
	//	查询用户信息
	repoUser, respUserErr := sev.UserRepo.Get(user)
	if respUserErr != nil {
		return nil, nil, errors.New("用户不存在")
	}
	//	查询用户权限信息
	repoRole, repoRoleErr := sev.RoleRepo.GetUseRoleByUserId(*repoUser)
	if repoRoleErr != nil {
		return nil, nil, errors.New("用户不存在")
	}
	return repoUser, repoRole, respUserErr
}
