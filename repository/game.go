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
		CreateGame(game *domain.GamePayload) *domain.Response
		GetLeaderboard(user_id string) *domain.Leaderboard
		AddLeaderboard(game *domain.GamePayload) *domain.Response
		UpdateLeaderboard(id int, game *domain.GamePayload) *domain.Response
		GetLeaderboards() *domain.Response
		GetHistoryGame(user_id string) *domain.Response
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

	if err := repo.db.Model(&domain.Question{}).Order("random()").Limit(5).Find(&questions).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(questions, 0, nil)
}

func (repo *gameRepository) CreateGame(game *domain.GamePayload) *domain.Response {
	if err := repo.db.Table("games").Create(&game).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(game, 0, nil)
}

func (repo *gameRepository) GetLeaderboard(user_id string) *domain.Leaderboard {
	leaderboard := new(domain.Leaderboard)

	if err := repo.db.Table("leaderboards").Where("user_id = ?", user_id).First(&leaderboard).Error; err != nil {
		return nil
	}

	return leaderboard
}

func (repo *gameRepository) AddLeaderboard(game *domain.GamePayload) *domain.Response {
	if err := repo.db.Table("leaderboards").Create(map[string]interface{}{
		"accumulated_score": game.Score,
		"user_id":           game.UserId,
	}).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(game, 0, nil)
}

func (repo *gameRepository) UpdateLeaderboard(id int, game *domain.GamePayload) *domain.Response {
	if err := repo.db.Table("leaderboards").Where("id = ?", id).UpdateColumn("accumulated_score", gorm.Expr("accumulated_score + ?", game.Score)).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(game, 0, nil)
}

func (repo *gameRepository) GetLeaderboards() *domain.Response {
	leaderboards := new(domain.Leaderboards)

	if err := repo.db.Table("leaderboards").Preload("User").Order("accumulated_score DESC").Find(&leaderboards).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(leaderboards, 0, nil)
}

func (repo *gameRepository) GetHistoryGame(user_id string) *domain.Response {
	games := new(domain.Games)

	if err := repo.db.Table("games").Where("user_id = ?", user_id).Limit(10).Find(&games).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(games, 0, nil)
}
