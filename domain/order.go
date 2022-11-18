package domain

import "time"

type Order struct {
	Id        int       `gorm:"primaryKey;column:id"`
	UserId    string    `gorm:"column:user_id"`
	CouponId  string    `gorm:"column:coupon_id"`
	Notes     string    `gorm:"column:notes"`
	PaymentId int       `gorm:"column:payment_id"`
	Status    string    `gorm:"column:status"`
	OrderDate time.Time `gorm:"column:order_date"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Orders []Order
