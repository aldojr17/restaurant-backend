package middleware

import (
	"final-project-backend/handler"
	"final-project-backend/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Contains(ctx.FullPath(), "docs") {
			ctx.Next()
			return
		}

		handler.ErrorResponse(ctx, http.StatusNotFound, util.ErrNoRoute)
	}
}
