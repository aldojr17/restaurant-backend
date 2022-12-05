package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	GameRepository interface {
		GetAllQuestions() *domain.Response
	}

	gameRepository struct {
		db *gorm.DB
	}
)

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepository{
		db: db,
	}
}

func (repo *gameRepository) GetAllQuestions() *domain.Response {
	questions := new(domain.Questions)

	if err := repo.db.Model(&domain.Question{}).Find(&questions).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(questions, 0, nil)
}
