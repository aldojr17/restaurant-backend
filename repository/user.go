package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		GetUserById(id string) *domain.Response
		CreateUser(user *domain.User) *domain.Response
		GetUserByEmail(email string) *domain.Response
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) GetUserById(id string) *domain.Response {
	user := new(domain.User)

	if err := repo.db.Where("id", id).First(&user).Error; err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return util.SetResponse(user, 0, nil)
}

func (repo *userRepository) GetUserByEmail(email string) *domain.Response {
	user := new(domain.User)

	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(user, 0, nil)
}

func (repo *userRepository) CreateUser(user *domain.User) *domain.Response {
	if err := repo.db.Create(&user).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseUserCreated, 0, nil)
}
