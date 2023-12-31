package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Review struct {
	Id          int       `gorm:"primaryKey;column:id" json:"id"`
	UserId      string    `gorm:"column:user_id" json:"user_id"`
	MenuId      int       `gorm:"column:menu_id" json:"menu_id"`
	Description string    `gorm:"column:description" json:"description"`
	Rating      int       `gorm:"column:rating" json:"rating"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	User        User      `gorm:"foreignKey:UserId;references:Id" json:"user"`
}

type Reviews []Review

func (r *Review) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}

	if r.MenuId <= 0 {
		return ErrMenuIdRequired
	}

	if r.Rating < 0 || r.Rating > 5 {
		return ErrRatingMinMax
	}

	return nil
}
