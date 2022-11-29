package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	UserRepository interface {
		GetUserById(id string) *domain.Response
		GetUserByEmail(email string) *domain.Response
		GetMenuFavorite(payload *domain.UserFavorite) *domain.Response

		CreateUser(user *domain.User) *domain.Response
		AddMenuFavorite(payload *domain.UserFavorite) *domain.Response

		UpdateUserData(id string, data map[string]interface{}) *domain.Response

		DeleteMenuFavorite(payload *domain.UserFavorite) *domain.Response
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
	user := new(domain.UserResponse)

	if err := repo.db.Preload("Favorites.Menu.Category").Preload("Favorites.Menu.MenuOption").Table("users").Where("id", id).First(&user).Error; err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return util.SetResponse(user, 0, nil)
}

func (repo *userRepository) GetUserByEmail(email string) *domain.Response {
	user := new(domain.User)

	if err := repo.db.Preload("Favorites").Where("email = ?", email).First(&user).Error; err != nil {
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

func (repo *userRepository) UpdateUserData(id string, data map[string]interface{}) *domain.Response {
	user := new(domain.UserResponse)

	if err := repo.db.Table("users").Model(&user).Clauses(clause.Returning{}).Where("id", id).Updates(data).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(user, 0, nil)
}

func (repo *userRepository) AddMenuFavorite(payload *domain.UserFavorite) *domain.Response {
	if err := repo.db.Create(&payload).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, domain.ErrMenuAlreadyAdded)
	}

	return util.SetResponse(domain.ResponseAddedToFavorite, 0, nil)
}

func (repo *userRepository) GetMenuFavorite(payload *domain.UserFavorite) *domain.Response {
	if err := repo.db.Where("user_id = ?", payload.UserId).Where("menu_id = ?", payload.MenuId).First(&payload).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, domain.ErrFavoriteMenuNotFound)
	}

	return util.SetResponse(payload, 0, nil)
}

func (repo *userRepository) DeleteMenuFavorite(payload *domain.UserFavorite) *domain.Response {
	if err := repo.db.Where("user_id = ?", payload.UserId).Where("menu_id = ?", payload.MenuId).Delete(&payload).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, domain.ErrMenuAlreadyAdded)
	}

	return util.SetResponse(domain.ResponseDeletedFromFavorite, 0, nil)
}
