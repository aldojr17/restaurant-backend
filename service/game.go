package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"time"

	"gorm.io/gorm"
)

type (
	GameService interface {
		GetAllQuestions() *domain.Response
		CreateGame(game *domain.GamePayload) *domain.Response
		GetLeaderboards() *domain.Response
		GetHistoryGame(user_id string) *domain.Response
	}

	gameService struct {
		db         *gorm.DB
		repo       repository.GameRepository
		couponRepo repository.CouponRepository
	}
)

func NewGameService(db *gorm.DB, repo repository.GameRepository, couponRepo repository.CouponRepository) GameService {
	return &gameService{
		db:         db,
		repo:       repo,
		couponRepo: couponRepo,
	}
}

func (s *gameService) GetAllQuestions() *domain.Response {
	return s.repo.GetAllQuestions()
}

func (s *gameService) GetLeaderboards() *domain.Response {
	return s.repo.GetLeaderboards()
}

func (s *gameService) GetHistoryGame(user_id string) *domain.Response {
	return s.repo.GetHistoryGame(user_id)
}

func (s *gameService) CreateGame(game *domain.GamePayload) *domain.Response {
	coupon := s.couponRepo.GetRandomCoupon()

	if game.Score >= 80 {
		if response := s.couponRepo.CheckUserCoupon(game.UserId, coupon.Id).Data; response != nil {
			s.couponRepo.UpdateCouponUser(game.UserId, coupon.Id)
		} else {
			valid, _ := time.Parse("2006-01-02T00:00:00Z", coupon.ValidUntil)
			s.couponRepo.AddCouponToUser(&domain.UserCoupon{
				UserId:    game.UserId,
				CouponId:  coupon.Id,
				ExpiredAt: valid,
				Qty:       1,
			})
		}
	}

	if data := s.repo.GetLeaderboard(game.UserId); data != nil {
		s.repo.UpdateLeaderboard(data.Id, game)
	} else {
		s.repo.AddLeaderboard(game)
	}

	return s.repo.CreateGame(game)
}
