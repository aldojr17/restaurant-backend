package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	CouponRepository interface {
		GetCouponOwnedByUser(user_id string) *domain.Response
		CreateCoupon(coupon *domain.Coupon) *domain.Response
		DeleteCoupon(id string) *domain.Response
	}

	couponRepository struct {
		db *gorm.DB
	}
)

func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &couponRepository{
		db: db,
	}
}

func (repo *couponRepository) GetCouponOwnedByUser(user_id string) *domain.Response {
	coupons := new(domain.UserCoupons)

	if err := repo.db.Preload("Coupon").Where("user_id", user_id).Where("expired_at > ?", time.Now()).Find(&coupons).Error; err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return util.SetResponse(coupons, 0, nil)
}

func (repo *couponRepository) CreateCoupon(coupon *domain.Coupon) *domain.Response {
	if err := repo.db.Create(&coupon).Error; err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return util.SetResponse(domain.ResponseCouponCreated, 0, nil)
}

func (repo *couponRepository) DeleteCoupon(id string) *domain.Response {
	coupon := new(domain.Coupon)

	if err := repo.db.Clauses(clause.Returning{}).Where("id", id).Delete(coupon).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(coupon, 0, nil)
}
