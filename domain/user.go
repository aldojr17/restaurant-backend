package domain

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
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

type UserProfile struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}

func (a *UserProfile) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(a); err != nil {
		return err
	}

	return nil
}
