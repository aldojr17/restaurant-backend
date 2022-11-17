package domain

import (
	"time"
)

type Coupon struct {
	Id        string    `gorm:"primaryKey;column:id"`
	Code      string    `gorm:"column:code"`
	Discount  int       `gorm:"column:discount"`
	CreatedAt time.Time `gorm:"column:created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

type UserCoupon struct {
	UserId    string    `gorm:"column:user_id"`
	CouponId  string    `gorm:"column:coupon_id"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	Coupon    Coupon    `gorm:"foreignKey:CouponId;references:Id"`
}

type UserCoupons []UserCoupon
