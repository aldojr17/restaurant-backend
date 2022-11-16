package repository

import (
	"database/sql"
	"final-project-backend/domain"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
	user *domain.User
}

func SetupSuite() *Suite {
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		log.Fatalf("mock db is null")
	}

	if s.mock == nil {
		log.Fatalf("sqlmock is null")
	}

	s.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return s
}
