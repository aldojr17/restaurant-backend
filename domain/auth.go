package domain

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AuthPayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(a); err != nil {
		return err
	}

	if a.Email == "" {
		return ErrEmailRequired
	}

	if strings.TrimSpace(a.Password) == "" {
		return ErrPasswordRequired
	}

	if !regex.MatchString(a.Email) {
		return ErrIncorrectEmailFormat
	}

	return nil
}

func (a *RegisterPayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(a); err != nil {
		return err
	}

	if a.FullName == "" {
		return ErrFullNameRequired
	}

	if a.Email == "" {
		return ErrEmailRequired
	}

	if strings.TrimSpace(a.Password) == "" {
		return ErrPasswordRequired
	}

	if !regex.MatchString(a.Email) {
		return ErrIncorrectEmailFormat
	}

	return nil
}
