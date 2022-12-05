package router

import (
	"final-project-backend/handler"
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func GameRoutes(router *gin.Engine, app *initialize.Application) {
	gameHandler := handler.NewGameHandler(app)
	router.GET("/questions", handler.GinHandlerWrapper(gameHandler.GetAllQuestions))
	router.GET("/game", handler.GinHandlerWrapper(gameHandler.GetHistoryGame))
	router.POST("/game", handler.GinHandlerWrapper(gameHandler.CreateGame))
	router.GET("/leaderboards", handler.GinHandlerWrapper(gameHandler.GetLeaderboards))
}
