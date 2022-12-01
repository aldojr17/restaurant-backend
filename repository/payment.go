package repository

import (
	"final-project-backend/domain"
	"final-project-backend/util"
	"net/http"

	"gorm.io/gorm"
)

type (
	PaymentRepository interface {
		GetAllPayment() *domain.Response
	}

	paymentRepository struct {
		db *gorm.DB
	}
)

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (repo *paymentRepository) GetAllPayment() *domain.Response {
	payments := new(domain.Payments)

	if err := repo.db.Find(&payments).Error; err != nil {
		return util.SetResponse(nil, http.StatusInternalServerError, err)
	}

	return util.SetResponse(payments, 0, nil)
}
