package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"
	"final-project-backend/util/jwt"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type (
	AuthService interface {
		Register(payload *domain.RegisterPayload) *domain.Response
		Login(payload *domain.AuthPayload) *domain.Response
	}

	authService struct {
		db       *gorm.DB
		userRepo repository.UserRepository
	}
)

func NewAuthService(db *gorm.DB, userRepo repository.UserRepository) AuthService {
	return &authService{
		db:       db,
		userRepo: userRepo,
	}
}

func (s *authService) Register(payload *domain.RegisterPayload) *domain.Response {
	hashedPass, err := util.GeneratePassword(payload.Password)
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	uuid := util.GenerateUUID()
	user := util.SetUser(uuid, payload.Email, hashedPass, payload.FullName)

	createUser := s.userRepo.CreateUser(user)
	if createUser.Err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrEmailAlreadyExists)
	}

	return util.SetResponse(domain.ResponseUserRegistered, 0, nil)
}

func (s *authService) Login(payload *domain.AuthPayload) *domain.Response {
	responseUser := s.userRepo.GetUserByEmail(payload.Email)
	if responseUser.Err != nil {
		return util.SetResponse(nil, http.StatusUnauthorized, domain.ErrWrongLoginCredential)
	}

	user := responseUser.Data.(*domain.User)

	if !util.ComparePassword(user.Password, payload.Password) {
		return util.SetResponse(nil, http.StatusBadRequest, domain.ErrWrongLoginCredential)
	}

	signedToken, err := jwt.GenerateToken(user.Id, user.Role)
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	fmt.Println("test")
	tokenResponse := domain.LoginResponse{
		Token: signedToken,
		User: domain.UserResponse{
			Id:             user.Id,
			Email:          user.Email,
			Address:        user.Address,
			FullName:       user.FullName,
			Phone:          user.Phone,
			ProfilePicture: user.ProfilePicture,
			Role:           user.Role,
		},
	}

	return util.SetResponse(tokenResponse, 0, nil)
}
