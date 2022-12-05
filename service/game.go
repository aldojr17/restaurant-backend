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

	return s.repo.CreateGame(game)
}
