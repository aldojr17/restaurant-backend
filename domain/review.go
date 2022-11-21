package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Review struct {
	Id          int       `gorm:"primaryKey;column:id"`
	UserId      string    `gorm:"column:user_id"`
	MenuId      int       `gorm:"column:menu_id" json:"menu_id"`
	Description string    `gorm:"column:description"`
	Rating      float32   `gorm:"column:rating"`
	CreatedAt   time.Time `gorm:"column:created_at"`
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
