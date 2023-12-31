package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"

	"gorm.io/gorm"
)

type (
	CouponService interface {
		CreateCoupon(coupon *domain.Coupon) *domain.Response
		DeleteCoupon(id string) *domain.Response
		GetAllCoupon() *domain.Response
		GetCoupon(id string) *domain.Response
		UpdateCoupon(coupon *domain.Coupon) *domain.Response
	}

	couponService struct {
		db   *gorm.DB
		repo repository.CouponRepository
	}
)

func NewCouponService(db *gorm.DB, repo repository.CouponRepository) CouponService {
	return &couponService{
		db:   db,
		repo: repo,
	}
}

func (s *couponService) CreateCoupon(coupon *domain.Coupon) *domain.Response {
	uuid := util.GenerateUUID()
	coupon.Id = uuid

	return s.repo.CreateCoupon(coupon)
}

func (s *couponService) DeleteCoupon(id string) *domain.Response {
	return s.repo.DeleteCoupon(id)
}

func (s *couponService) GetAllCoupon() *domain.Response {
	return s.repo.GetAllCoupon()
}

func (s *couponService) GetCoupon(id string) *domain.Response {
	return s.repo.GetCoupon(id)
}

func (s *couponService) UpdateCoupon(coupon *domain.Coupon) *domain.Response {
	return s.repo.UpdateCoupon(coupon)
}
