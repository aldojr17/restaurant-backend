package domain

import (
	"time"
)

type Coupon struct {
	Id        string    `gorm:"primaryKey;column:id" json:"id"`
	Code      string    `gorm:"column:code" json:"code"`
	Discount  int       `gorm:"column:discount" json:"discount"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type UserCoupon struct {
	UserId    string    `gorm:"column:user_id" json:"user_id"`
	CouponId  string    `gorm:"column:coupon_id" json:"coupon_id"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
	Coupon    Coupon    `gorm:"foreignKey:CouponId;references:Id" json:"coupon"`
}

type UserCoupons []UserCoupon
