package domain

import (
	"errors"
	"regexp"
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrPasswordRequired = errors.New("password is required")

	ErrIncorrectEmailFormat = errors.New("invalid email format")
)

const (
	EMAIL = "email"

	ResponseUserRegistered = "User Registered."
	ResponseUserCreated    = "User Created."
)

var regex = regexp.MustCompile("^[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*$")
