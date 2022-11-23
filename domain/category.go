package domain

type Category struct {
	Id   int    `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type Categories []Category
