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
	ErrWalletNotFound         = errorNotFound("wallet")
	ErrNoRoute                = errorNotFound("route")
	ErrUserNotFound           = errorNotFound("user")
	ErrWrongLoginCredential   = NewError("wrong email or password")
	ErrInvalidTopUpAmount     = NewError("invalid amount, should between 50000 and 10000000")
	ErrInvalidTransferAmount  = NewError("invalid amount, should between 1000 and 50000000")
	ErrInvalidSourceOfFunds   = NewError("invalid source of funds, should between 1 and 3")
	ErrTokenInvalid           = NewError("token invalid")
	ErrNoAuthorization        = NewError("no authorization header provided")
	ErrDescriptionLimitExceed = NewError("description length limit is 35 characters")
	ErrInsufficientBalance    = NewError("insufficient balance")
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
