package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Question struct {
	Id            int    `gorm:"primaryKey;column:id" json:"id"`
	Question      string `gorm:"column:question" json:"question"`
	CorrectAnswer string `gorm:"column:correct_answer" json:"correct_answer"`
	OptionOne     string `gorm:"column:option_one" json:"option_one"`
	OptionTwo     string `gorm:"column:option_two" json:"option_two"`
	OptionThree   string `gorm:"column:option_three" json:"option_three"`
	OptionFour    string `gorm:"column:option_four" json:"option_four"`
}

type GamePayload struct {
	UserId    string    `gorm:"column:user_id" json:"user_id"`
	Score     int       `gorm:"column:score" json:"score"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type Leaderboard struct {
	Id               int    `gorm:"primaryKey;column:id" json:"id"`
	UserId           string `gorm:"column:user_id" json:"user_id"`
	AccumulatedScore int    `gorm:"column:accumulated_score" json:"accumulated_score"`
	User             User   `gorm:"foreignKey:UserId;references:Id" json:"user"`
}

type Games []GamePayload

type Leaderboards []Leaderboard

type Questions []Question

func (game *GamePayload) Validate(c *gin.Context) error {
	if err := c.ShouldBindJSON(game); err != nil {
		return err
	}

	return nil
}
