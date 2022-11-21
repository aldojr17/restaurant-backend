package domain

import (
	"errors"
	"regexp"
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required")
	ErrMenuIdRequired   = errors.New("menu id is required")

	ErrIncorrectEmailFormat = errors.New("invalid email format")
	ErrWrongLoginCredential = errors.New("wrong email or password")
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrMenuAlreadyAdded     = errors.New("menu already added to favorite")
)

const (
	USER_ID = "user_id"
	ROLE_ID = "role_id"

	ResponseUserRegistered     = "User Registered."
	ResponseUserCreated        = "User Created."
	ResponseReviewAdded        = "Review Added."
	ResponseAddedToFavorite    = "Menu added to favorite."
	ResponseOrderStatusUpdated = "Order Status Updated."
	ResponseCouponCreated      = "Coupon Created."

	DELIVERY_STATUS_IN_PROGRESS = "In Progress"
	DELIVERY_STATUS_IN_TRANSIT  = "In Transit"
	DELIVERY_STATUS_RECEIVED    = "Received"
)

var regex = regexp.MustCompile("^[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*$")
