package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	CouponRepository interface {
		GetCouponOwnedByUser(user_id string) *domain.Response
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

	if err := repo.db.Preload("Coupon").Table("user_coupons").Where("user_id", user_id).Find(&coupons).Error; err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return util.SetResponse(coupons, 0, nil)
}
