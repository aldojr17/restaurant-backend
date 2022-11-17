package repository

import (
	"errors"
	"final-project-backend/util"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCouponOwnedByUser(t *testing.T) {
	s := SetupSuite()

	uuid := util.GenerateUUID()

	rowsPreload := s.mock.NewRows([]string{"Id", "Code", "Discount", "CreatedAt", "DeletedAt"}).
		AddRow(uuid, "TEST01", 10000, time.Now(), nil)
	rows := s.mock.NewRows([]string{"UserId", "CouponId", "ExpiredAt"}).
		AddRow(uuid, uuid, time.Now())

	queryPreload := `SELECT * FROM "coupons" WHERE "coupons"."id" = $1`
	query := `SELECT * FROM "user_coupons" WHERE "user_id" = $1`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)
	s.mock.ExpectQuery(regexp.QuoteMeta(queryPreload)).WillReturnRows(rowsPreload)

	repo := NewCouponRepository(s.db)

	response := repo.GetCouponOwnedByUser(uuid)
	if response.Err != nil {
		t.Errorf("Failed to select user by id, got error: %v", response.Err)
		t.FailNow()
	}

	assert.NotNil(t, response.Data)
}

func TestCouponOwnedByUserError(t *testing.T) {
	s := SetupSuite()

	uuid := util.GenerateUUID()

	rowsPreload := s.mock.NewRows([]string{"Id", "Code", "Discount", "CreatedAt", "DeletedAt"}).
		AddRow(uuid, "TEST01", 10000, time.Now(), nil)

	queryPreload := `SELECT * FROM "coupons" WHERE "coupons"."id" = $1`
	query := `SELECT * FROM "user_coupons" WHERE "user_id" = $1`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New(("error")))
	s.mock.ExpectQuery(regexp.QuoteMeta(queryPreload)).WillReturnRows(rowsPreload)

	repo := NewCouponRepository(s.db)

	response := repo.GetCouponOwnedByUser(uuid)
	if response.Err == nil {
		t.Errorf("Failed to select user by id, got error: %v", response.Err)
		t.FailNow()
	}

	assert.NotNil(t, response.Err)
}
