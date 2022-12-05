package domain

type Question struct {
	Id            int    `gorm:"primaryKey;column:id" json:"id"`
	Question      string `gorm:"column:question" json:"question"`
	CorrectAnswer string `gorm:"column:correct_answer" json:"correct_answer"`
	OptionOne     string `gorm:"column:option_one" json:"option_one"`
	OptionTwo     string `gorm:"column:option_two" json:"option_two"`
	OptionThree   string `gorm:"column:option_three" json:"option_three"`
	OptionFour    string `gorm:"column:option_four" json:"option_four"`
}

type Questions []Question
