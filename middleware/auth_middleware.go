package middleware

import (
	"final-project-backend/domain"
	"final-project-backend/handler"
	"final-project-backend/util"
	"final-project-backend/util/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Contains(ctx.FullPath(), "docs") || strings.Contains(ctx.FullPath(), "login") || strings.Contains(ctx.FullPath(), "register") || (strings.Contains(ctx.FullPath(), "menus") && ctx.Request.Method == "GET") || strings.Contains(ctx.FullPath(), "categories") || strings.Contains(ctx.FullPath(), "payments") {
			ctx.Next()
			return
		}

		authToken := ctx.Request.Header.Get("Authorization")
		if authToken == "" {
			handler.ErrorResponse(ctx, http.StatusUnauthorized, util.ErrUnauthorized)
			ctx.Abort()
			return
		}

		authToken = strings.Replace(authToken, "Bearer ", "", 1)

		user_id, role_id, err := jwt.ValidateToken(authToken)
		if err != nil || user_id == "" {
			handler.ErrorResponse(ctx, http.StatusUnauthorized, util.ErrUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set(domain.USER_ID, user_id)
		ctx.Set(domain.ROLE_ID, role_id)

		ctx.Next()
	}
}
