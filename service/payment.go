package service

import (
	"final-project-backend/domain"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type (
	PaymentService interface {
		GetAllPayment() *domain.Response
	}

	paymentService struct {
		db   *gorm.DB
		repo repository.PaymentRepository
	}
)

func NewPaymentService(db *gorm.DB, repo repository.PaymentRepository) PaymentService {
	return &paymentService{
		db:   db,
		repo: repo,
	}
}

func (s *paymentService) GetAllPayment() *domain.Response {
	return s.repo.GetAllPayment()
}
