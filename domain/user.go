package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id             string         `gorm:"primaryKey;column:id" json:"id"`
	Email          string         `gorm:"column:email" json:"email"`
	Password       string         `gorm:"column:password" json:"-"`
	Address        *string        `gorm:"column:address" json:"address"`
	FullName       string         `gorm:"column:full_name" json:"full_name"`
	Phone          *string        `gorm:"column:phone" json:"phone"`
	ProfilePicture *string        `gorm:"column:profile_picture" json:"profile"`
	Role           int            `gorm:"column:role" json:"-"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"-"`
	Favorites      []UserFavorite `gorm:"foreignKey:UserId;references:Id" json:"favorites"`
}

type UserProfile struct {
	UserId         string `json:"user_id"`
	Address        string `json:"address"`
	FullName       string `json:"full_name"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}

type UserFavorite struct {
	UserId string `gorm:"column:user_id" json:"user_id"`
	MenuId int    `gorm:"column:menu_id" json:"menu_id"`
	Menu   Menu   `gorm:"foreignKey:MenuId;references:Id" json:"menu"`
}

func (u *UserProfile) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(u); err != nil {
		return err
	}

	return nil
}

func (u *UserFavorite) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(u); err != nil {
		return err
	}

	if u.MenuId <= 0 {
		return ErrMenuIdRequired
	}

	return nil
}
