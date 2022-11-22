package domain

type Category struct {
	Id   int    `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type CategoryWithMenu struct {
	Id    int    `gorm:"primaryKey;column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Menus []Menu `gorm:"foreignKey:CategoryId;references:Id" json:"menus"`
}

type Categories []CategoryWithMenu
