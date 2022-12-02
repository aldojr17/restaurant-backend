package domain

import (
	"errors"
	"regexp"
)

var (
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrFullNameRequired   = errors.New("fullname is required")
	ErrMenuIdRequired     = errors.New("menu id is required")
	ErrPaymentIdRequired  = errors.New("payment id is required")
	ErrOrdersRequired     = errors.New("orders is required")
	ErrTotalPriceRequired = errors.New("total price is required")
	ErrOrderIdRequired    = errors.New("order id is required")
	ErrQtyRequired        = errors.New("qty is required")

	ErrIncorrectEmailFormat = errors.New("invalid email format")
	ErrWrongLoginCredential = errors.New("wrong email or password")
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrMenuAlreadyAdded     = errors.New("menu already added to favorite")
	ErrMenuNotFound         = errors.New("menu not found")
	ErrFavoriteMenuNotFound = errors.New("favorite menu not found")
	ErrCouponInvalid        = errors.New("coupon invalid")
	ErrRatingMinMax         = errors.New("rating min 0 and max 5")
	ErrInvalidStatus        = errors.New("invalid order status (Preparing, On The Way, Received)")
)

const (
	USER_ID = "user_id"
	ROLE_ID = "role_id"

	ResponseUserRegistered      = "User Registered."
	ResponseUserCreated         = "User Created."
	ResponseReviewAdded         = "Review Added."
	ResponseReviewUpdated       = "Review Updated."
	ResponseMenuUpdated         = "Menu Updated."
	ResponseAddedToFavorite     = "Menu added to favorite."
	ResponseDeletedFromFavorite = "Menu deleted from favorite."
	ResponseOrderCreated        = "Order Created."
	ResponseOrderDetailsCreated = "Order Details Created."
	ResponseOrderStatusUpdated  = "Order Status Updated."
	ResponseCouponCreated       = "Coupon Created."

	DELIVERY_STATUS_PREPARING  = "Preparing"
	DELIVERY_STATUS_ON_THE_WAY = "On The Way"
	DELIVERY_STATUS_RECEIVED   = "Received"
)

var regex = regexp.MustCompile("^[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*$")
