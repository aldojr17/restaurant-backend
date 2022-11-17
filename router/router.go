package router

import (
	"final-project-backend/initialize"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, app *initialize.Application) {
	AuthRoutes(router, app)
	UserRoutes(router, app)
	MenuRoutes(router, app)
}
