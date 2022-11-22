package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Menu struct {
	Id          int            `gorm:"primaryKey;column:id" json:"id"`
	Name        string         `gorm:"column:name" json:"name"`
	Price       int            `gorm:"column:price" json:"price"`
	Photo       string         `gorm:"column:photo" json:"photo"`
	CategoryId  int            `gorm:"column:category_id" json:"category_id"`
	Rating      *float32       `gorm:"column:rating" json:"rating"`
	TotalReview *int           `gorm:"column:total_review" json:"total_review"`
	IsAvailable bool           `gorm:"column:is_available" json:"is_available"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Category    Category       `gorm:"foreignKey:CategoryId;references:Id" json:"category"`
}

type MenuPayload struct {
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Photo       string  `json:"photo"`
	CategoryId  int     `json:"category_id"`
	Rating      float32 `json:"rating"`
	IsAvailable bool    `json:"is_available"`
}

type MenuOption struct {
	MenuId     int    `gorm:"primaryKey;column:menu_id" json:"menu_id"`
	OrderId    int    `gorm:"primaryKey;column:order_id" json:"order_id"`
	Qty        int    `gorm:"qty" json:"qty"`
	Options    string `gorm:"options" json:"options"`
	MenuDetail Menu   `gorm:"foreignKey:MenuId;references:Id" json:"menu_detail"`
}

type Menus []Menu

func (m *MenuPayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(m); err != nil {
		return err
	}

	return nil
}
