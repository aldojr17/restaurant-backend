package domain

import (
	"time"
)

type Menu struct {
	Id          string    `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Price       int       `gorm:"column:price" json:"price"`
	Photo       string    `gorm:"column:photo" json:"photo"`
	CategoryId  int       `gorm:"column:category_id" json:"category_id"`
	Rating      float32   `gorm:"column:rating" json:"rating"`
	IsAvailable bool      `gorm:"column:is_available" json:"is_available"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	Category    Category  `gorm:"foreignKey:CategoryId;references:Id" json:"category"`
}

type MenuOption struct {
	MenuId  int    `gorm:"primaryKey;column:menu_id"`
	OrderId int    `gorm:"primaryKey;column:order_id"`
	Qty     int    `gorm:"qty"`
	Options string `gorm:"options"`
}

type Menus []Menu
