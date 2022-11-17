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
	Address        sql.NullString `gorm:"column:address"`
	FullName       sql.NullString `gorm:"column:full_name"`
	Phone          sql.NullString `gorm:"column:phone"`
	ProfilePicture sql.NullString `gorm:"column:profile_picture"`
	Role           int            `gorm:"column:role"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
}

type UserProfile struct {
	UserId         string `json:"user_id"`
	Address        string `json:"address"`
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
