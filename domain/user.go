package domain

import (
	"database/sql"
	"time"
)

const (
	EMAIL               = "email"
	ResponseUserCreated = "User Created."
)

type User struct {
	Id             string         `gorm:"primaryKey;column:id"`
	Email          string         `gorm:"column:email"`
	Password       string         `gorm:"column:password"`
	Username       sql.NullString `gorm:"column:username"`
	FullName       sql.NullString `gorm:"column:full_name"`
	Phone          sql.NullString `gorm:"column:phone"`
	ProfilePicture sql.NullString `gorm:"column:profile_picture"`
	Role           int            `gorm:"column:role"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
}
