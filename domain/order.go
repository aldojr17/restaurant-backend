package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Order struct {
	Id          int          `gorm:"primaryKey;column:id" json:"id"`
	UserId      string       `gorm:"column:user_id" json:"user_id"`
	CouponId    *string      `gorm:"column:coupon_id" json:"coupon_id"`
	Notes       *string      `gorm:"column:notes" json:"notes"`
	PaymentId   int          `gorm:"column:payment_id" json:"payment_id"`
	Status      string       `gorm:"column:status" json:"status"`
	OrderDate   time.Time    `gorm:"column:order_date" json:"order_date"`
	UpdatedAt   time.Time    `gorm:"column:updated_at" json:"updated_at"`
	MenuOptions []MenuOption `gorm:"foreignKey:OrderId;references:Id" json:"menu_options"`
	Payment     Payment      `gorm:"foreignKey:PaymentId;references:Id" json:"payment_detail"`
}

type OrderPayload struct {
	UserId    string
	CouponId  *string   `json:"coupon_id"`
	Notes     *string   `json:"notes"`
	PaymentId int       `json:"payment_id"`
	Status    string    `json:"status"`
	OrderDate time.Time `json:"order_date"`
}

type OrderStatusPayload struct {
	Id     int    `gorm:"-" json:"id"`
	Status string `gorm:"column:status" json:"status"`
}

type Orders []Order

func (o *OrderStatusPayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(o); err != nil {
		return err
	}

	return nil
}

func (o *OrderPayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(o); err != nil {
		return err
	}

	if o.PaymentId <= 0 {
		return ErrPaymentIdRequired
	}

	if o.Status == "" {
		o.Status = DELIVERY_STATUS_IN_PROGRESS
	}

	if o.Status != DELIVERY_STATUS_IN_PROGRESS && o.Status != DELIVERY_STATUS_IN_TRANSIT && o.Status != DELIVERY_STATUS_RECEIVED {
		return ErrInvalidStatus
	}

	if o.CouponId != nil && *o.CouponId == "" {
		o.CouponId = nil
	}

	o.OrderDate = time.Now()

	return nil
}
