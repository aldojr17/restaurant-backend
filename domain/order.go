package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Order struct {
	Id           int           `gorm:"primaryKey;column:id" json:"id"`
	UserId       string        `gorm:"column:user_id" json:"user_id"`
	CouponId     *string       `gorm:"column:coupon_id" json:"coupon_id"`
	PaymentId    int           `gorm:"column:payment_id" json:"payment_id"`
	Status       string        `gorm:"column:status" json:"status"`
	OrderDate    time.Time     `gorm:"column:order_date" json:"order_date"`
	TotalPrice   int           `gorm:"column:total_price" json:"total_price"`
	UpdatedAt    time.Time     `gorm:"column:updated_at" json:"updated_at"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderId;references:Id" json:"order_details"`
	Payment      Payment       `gorm:"foreignKey:PaymentId;references:Id" json:"payment_detail"`
	Subtotal     int           `gorm:"column:subtotal" json:"subtotal"`
	Coupon       Coupon        `gorm:"foreignKey:CouponId;references:Id" json:"coupon"`
	User         User          `gorm:"foreignKey:UserId;references:Id" json:"user"`
}

type OrderPayload struct {
	Id         int `json:"id"`
	UserId     string
	CouponId   *string              `json:"coupon_id"`
	PaymentId  int                  `json:"payment_id"`
	Status     string               `json:"status"`
	OrderDate  time.Time            `json:"order_date"`
	TotalPrice int                  `json:"total_price"`
	Subtotal   int                  `json:"subtotal"`
	Orders     []OrderDetailPayload `gorm:"-" json:"orders"`
}

type OrderStatusPayload struct {
	Id     int    `gorm:"-" json:"id"`
	Status string `gorm:"column:status" json:"status"`
}

type OrderDetail struct {
	Id         int  `gorm:"primaryKey;column:id" json:"id"`
	MenuId     int  `gorm:"column:menu_id" json:"menu_id"`
	OrderId    int  `gorm:"column:order_id" json:"order_id"`
	Qty        int  `gorm:"qty" json:"qty"`
	OptionId   int  `gorm:"option_id" json:"option_id"`
	MenuDetail Menu `gorm:"foreignKey:MenuId;references:Id" json:"menu_detail"`
}

type OrderDetailPayload struct {
	MenuId   int  `json:"menu_id"`
	OrderId  int  `json:"order_id"`
	Qty      int  `json:"qty"`
	OptionId *int `json:"option_id"`
}

type OrderFilter struct {
	Start time.Time
	End   time.Time
}

type Orders []Order

func (o *OrderStatusPayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(o); err != nil {
		return err
	}

	return nil
}

func (o *OrderDetailPayload) Validate(c *gin.Context) error {
	if o.MenuId == 0 {
		return ErrMenuIdRequired
	}

	if o.Qty == 0 {
		return ErrQtyRequired
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
		o.Status = DELIVERY_STATUS_PREPARING
	}

	if o.Status != DELIVERY_STATUS_ON_THE_WAY && o.Status != DELIVERY_STATUS_PREPARING && o.Status != DELIVERY_STATUS_RECEIVED {
		return ErrInvalidStatus
	}

	if o.CouponId != nil && *o.CouponId == "" {
		o.CouponId = nil
	}

	if len(o.Orders) == 0 {
		return ErrOrderDetailsRequired
	}

	for _, order := range o.Orders {
		if err := order.Validate(c); err != nil {
			return err
		}
	}

	o.OrderDate = time.Now()

	return nil
}
