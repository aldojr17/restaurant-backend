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
		GetValidCoupon(user_id string, coupon_id string) *domain.Response
		ReduceQty(user_id string, coupon_id string) error
		GetAllCoupon() *domain.Response
		GetCoupon(id string) *domain.Response
		UpdateCoupon(coupon *domain.Coupon) *domain.Response
		AddCouponToUser(coupon *domain.UserCoupon) *domain.Response
		UpdateCouponUser(user_id string, coupon_id string) *domain.Response
		GetRandomCoupon() *domain.Coupon
		CheckUserCoupon(user_id string, coupon_id string) *domain.Response
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

	if err := repo.db.Preload("Coupon").Where("user_id", user_id).Where("expired_at > ?", time.Now()).Where("qty > ?", 0).Find(&coupons).Error; err != nil {
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

func (repo *couponRepository) GetValidCoupon(user_id string, coupon_id string) *domain.Response {
	coupon := new(domain.UserCoupon)

	if err := repo.db.Where("user_id = ? AND coupon_id = ?", user_id, coupon_id).Where("expired_at > ?", time.Now()).
		Where("qty > ?", 0).First(&coupon).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(coupon, 0, nil)
}

func (repo *couponRepository) ReduceQty(user_id string, coupon_id string) error {
	if err := repo.db.Table("user_coupons").Where("user_id = ? AND coupon_id = ?", user_id, coupon_id).Where("expired_at > ?", time.Now()).UpdateColumn("qty", gorm.Expr("qty - 1")).Error; err != nil {
		return err
	}

	return nil
}

func (repo *couponRepository) GetAllCoupon() *domain.Response {
	coupons := new(domain.Coupons)

	if err := repo.db.Where("valid_until > ?", time.Now()).Order("created_at DESC").Find(&coupons).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(coupons, 0, nil)
}

func (repo *couponRepository) GetCoupon(id string) *domain.Response {
	coupon := new(domain.Coupon)

	if err := repo.db.Where("valid_until > ?", time.Now()).Where("id = ?", id).Find(&coupon).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(coupon, 0, nil)
}

func (repo *couponRepository) UpdateCoupon(coupon *domain.Coupon) *domain.Response {
	if err := repo.db.Model(&coupon).Where("id = ?", coupon.Id).Updates(map[string]interface{}{
		"discount":    coupon.Discount,
		"valid_until": coupon.ValidUntil,
	}).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(coupon, 0, nil)
}

func (repo *couponRepository) CheckUserCoupon(user_id string, coupon_id string) *domain.Response {
	coupon := new(domain.UserCoupon)

	if err := repo.db.Where("user_id = ? AND coupon_id = ?", user_id, coupon_id).First(&coupon).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(coupon, 0, nil)
}

func (repo *couponRepository) AddCouponToUser(coupon *domain.UserCoupon) *domain.Response {
	if err := repo.db.Create(&coupon).Error; err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, err)
	}

	return util.SetResponse(domain.ResponseCouponAdded, 0, nil)
}

func (repo *couponRepository) UpdateCouponUser(user_id string, coupon_id string) *domain.Response {
	if err := repo.db.Table("user_coupons").Where("user_id = ? AND coupon_id = ?", user_id, coupon_id).Where("expired_at > ?", time.Now()).UpdateColumn("qty", gorm.Expr("qty + 1")).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(domain.ResponseCouponAdded, 0, nil)
}

func (repo *couponRepository) GetRandomCoupon() *domain.Coupon {
	coupon := new(domain.Coupon)

	if err := repo.db.Table("coupons").Order("random()").Limit(1).First(&coupon).Error; err != nil {
		return nil
	}

	return coupon
}
