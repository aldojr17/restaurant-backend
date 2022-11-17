package domain

type Category struct {
	Id   int    `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`
}
