package domain

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
)

var (
	ErrEmailRequired        = errors.New("email is required")
	ErrIncorrectEmailFormat = errors.New("invalid email format")
	ErrPasswordRequired     = errors.New("password is required")
)

const (
	ResponseUserRegistered = "User Registered."
)

var regex = regexp.MustCompile("^[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+(?:.[a-zA-Z0-9]+)*$")

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (a *AuthPayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(a); err != nil {
		return err
	}

	if a.Email == "" {
		return ErrEmailRequired
	}

	if a.Password == "" {
		return ErrPasswordRequired
	}

	if !regex.MatchString(a.Email) {
		return ErrIncorrectEmailFormat
	}

	return nil
}
