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

	ResponseUserRegistered  = "User Registered."
	ResponseUserCreated     = "User Created."
	ResponseAddedToFavorite = "Menu added to favorite."
)

var regex = regexp.MustCompile("^[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*$")
