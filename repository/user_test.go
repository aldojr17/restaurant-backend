package repository

import (
	"errors"
	"final-project-backend/domain"
	"final-project-backend/util"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	s := SetupSuite()

	uuid := util.GenerateUUID()

	rows := s.mock.NewRows([]string{"Id", "Email", "Password", "Username", "FullName", "Phone", "ProfilePicture", "Role", "CreatedAt", "UpdatedAt"}).
		AddRow(uuid, "test_id", "1234", nil, nil, nil, nil, 1, time.Now(), time.Now())

	query := `SELECT * FROM "users" WHERE "id" = $1 ORDER BY "users"."id" LIMIT 1`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	repo := NewUserRepository(s.db)

	response := repo.GetUserById(uuid)
	if response.Err != nil {
		t.Errorf("Failed to select user by id, got error: %v", response.Err)
		t.FailNow()
	}

	assert.NotNil(t, response.Data)
}

func TestGetUserByIdError(t *testing.T) {
	s := SetupSuite()

	uuid := util.GenerateUUID()

	query := `SELECT * FROM "users" WHERE "id" = $1 ORDER BY "users"."id" LIMIT 1`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))

	repo := NewUserRepository(s.db)

	response := repo.GetUserById(uuid)
	if response.Err == nil {
		t.Errorf("Failed to select user by id, got error: %v", response.Err)
		t.FailNow()
	}

	assert.Nil(t, response.Data)
}

func TestGetUserByEmail(t *testing.T) {
	s := SetupSuite()

	uuid := util.GenerateUUID()

	rows := s.mock.NewRows([]string{"Id", "Email", "Password", "Username", "FullName", "Phone", "ProfilePicture", "Role", "CreatedAt", "UpdatedAt"}).
		AddRow(uuid, "test_id@mail.com", "1234", nil, nil, nil, nil, 1, time.Now(), time.Now())

	query := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	repo := NewUserRepository(s.db)

	response := repo.GetUserByEmail("test_id@mail.com")
	if response.Err != nil {
		t.Errorf("Failed to select user by email, got error: %v", response.Err)
		t.FailNow()
	}

	assert.NotNil(t, response.Data)
}

func TestGetUserByEmailError(t *testing.T) {
	s := SetupSuite()

	query := `SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("error"))

	repo := NewUserRepository(s.db)

	response := repo.GetUserByEmail("test_id")
	if response.Err == nil {
		t.Errorf("Failed to select user by email, got error: %v", response.Err)
		t.FailNow()
	}

	assert.Nil(t, response.Data)
}

func TestCreateUser(t *testing.T) {
	s := SetupSuite()

	now := time.Now()
	uuid := util.GenerateUUID()
	s.user = &domain.User{
		Id:        uuid,
		Email:     "test@gmail.com",
		Password:  "1234",
		Role:      1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	query := `INSERT INTO "users" ("id","email","password","username","full_name","phone","profile_picture","role","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(uuid, s.user.Email, s.user.Password, nil, nil, nil, nil, 1, now, now).WillReturnResult(sqlmock.NewResult(0, 0))
	s.mock.ExpectCommit()

	repo := NewUserRepository(s.db)

	response := repo.CreateUser(s.user)
	if response.Err != nil {
		t.Errorf("Failed to create user, got error: %+v", response.Err)
		t.FailNow()
	}
}

func TestCreateUserError(t *testing.T) {
	s := SetupSuite()

	now := time.Now()
	uuid := util.GenerateUUID()
	s.user = &domain.User{
		Id:        uuid,
		Email:     "test@gmail.com",
		Password:  "1234",
		CreatedAt: now,
	}

	query := `INSERT INTO "users" ("id","email","password","username","full_name","phone","profile_picture","role","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(uuid, s.user.Email, s.user.Password, nil, nil, nil, nil, 1, now, now).WillReturnError(errors.New("error"))
	s.mock.ExpectCommit()

	repo := NewUserRepository(s.db)

	response := repo.CreateUser(s.user)
	if response.Err == nil {
		t.Errorf("Failed to create user, got error: %+v", response.Err)
		t.FailNow()
	}
}
