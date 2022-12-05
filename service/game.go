package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type (
	GameService interface {
		GetAllQuestions() *domain.Response
	}

	gameService struct {
		db   *gorm.DB
		repo repository.GameRepository
	}
)

func NewGameService(db *gorm.DB, repo repository.GameRepository) GameService {
	return &gameService{
		db:   db,
		repo: repo,
	}
}

func (s *gameService) GetAllQuestions() *domain.Response {
	return s.repo.GetAllQuestions()
}
