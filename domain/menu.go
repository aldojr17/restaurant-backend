package domain

import (
	"database/sql"
	"time"
)

type Menu struct {
	Id          string          `gorm:"primaryKey;column:id"`
	Name        string          `gorm:"column:name"`
	Price       int             `gorm:"column:price"`
	Photo       string          `gorm:"column:photo"`
	CategoryId  int             `gorm:"column:category_id"`
	Rating      sql.NullFloat64 `gorm:"column:rating"`
	IsAvailable bool            `gorm:"column:is_available"`
	CreatedAt   time.Time       `gorm:"column:created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at"`
	Category    Category        `gorm:"foreignKey:CategoryId;references:Id"`
}

type Menus []Menu
