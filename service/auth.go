package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"
	"final-project-backend/util"
	"final-project-backend/util/jwt"
	"net/http"

	"gorm.io/gorm"
)

type (
	AuthService interface {
		Register(payload *domain.AuthPayload) *domain.Response
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

func (s *authService) Register(payload *domain.AuthPayload) *domain.Response {
	hashedPass, err := util.GeneratePassword(payload.Password)
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	uuid := util.GenerateUUID()
	user := util.SetUser(uuid, payload.Email, hashedPass)

	createUser := s.userRepo.CreateUser(user)
	if createUser.Err != nil {
		return util.SetResponse(nil, http.StatusBadRequest, createUser.Err)
	}

	return util.SetResponse(domain.ResponseUserRegistered, 0, nil)
}

func (s *authService) Login(payload *domain.AuthPayload) *domain.Response {
	responseUser := s.userRepo.GetUserByEmail(payload.Email)
	if responseUser.Err != nil {
		return util.SetResponse(nil, http.StatusUnauthorized, responseUser.Err)
	}

	user := responseUser.Data.(*domain.User)

	if !util.ComparePassword(user.Password, payload.Password) {
		return util.SetResponse(nil, http.StatusBadRequest, util.ErrWrongLoginCredential)
	}

	signedToken, err := jwt.GenerateToken(user.Id)
	if err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	tokenResponse := domain.LoginResponse{
		Token: signedToken,
	}

	return util.SetResponse(tokenResponse, 0, nil)
}
