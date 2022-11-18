package util

import (
	"errors"
	"final-project-backend/domain"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUnauthorized = NewError("unauthorized")

	ErrNoRoute = errorNotFound("route")
)

func GeneratePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(password, input string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(input)) == nil
}

func SetResponse(data interface{}, code int, err error) *domain.Response {
	return &domain.Response{
		Data: data,
		Code: code,
		Err:  err,
	}
}

func SetUser(id, email, password string) *domain.User {
	user := new(domain.User)
	user.Id = id
	user.Email = strings.ToLower(email)
	user.Password = password
	user.Role = 1
	return user
}

func GenerateUUID() string {
	return uuid.New().String()
}

func NewError(msg string) error {
	return errors.New(msg)
}

func errorNotFound(entity string) error {
	return fmt.Errorf("%s not found", entity)
}
