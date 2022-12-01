package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Coupon struct {
	Id         string         `gorm:"primaryKey;column:id" json:"id"`
	Code       string         `gorm:"column:code" json:"code"`
	Discount   int            `gorm:"column:discount" json:"discount"`
	ValidUntil string         `gorm:"column:valid_until" json:"valid_until"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

type UserCoupon struct {
	UserId    string    `gorm:"column:user_id" json:"user_id"`
	CouponId  string    `gorm:"column:coupon_id" json:"-"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
	Qty       int       `gorm:"column:qty" json:"qty"`
	Coupon    Coupon    `gorm:"foreignKey:CouponId;references:Id" json:"coupon"`
}

type UserCoupons []UserCoupon
type Coupons []Coupon

func (coupon *Coupon) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(coupon); err != nil {
		return err
	}

	return nil
}
