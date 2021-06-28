package repository

import (
	"go-service/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type IUserRepository interface {
	Get(user models.User) (*models.User, error)
	ExistById(user models.User) (*models.User, error)
}

//	查询用户
func (repo *UserRepository) Get(user models.User) (*models.User, error) {
	err := repo.DB.Where("userName = ? and userPass = ?", user.UserName, user.UserPass).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

//	Exist 查询是否存在用户
func (repo *UserRepository) ExistById(user models.User) (*models.User, error) {
	err := repo.DB.Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
