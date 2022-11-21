package domain

type Payment struct {
	Id          int    `gorm:"primaryKey;column:id" json:"id"`
	Description string `gorm:"column:description" json:"description"`
}
