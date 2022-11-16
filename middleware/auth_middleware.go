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
		if strings.Contains(ctx.FullPath(), "docs") || strings.Contains(ctx.FullPath(), "login") || strings.Contains(ctx.FullPath(), "register") {
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

		email, err := jwt.ValidateToken(authToken)
		if err != nil || email == "" {
			handler.ErrorResponse(ctx, http.StatusUnauthorized, util.ErrUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set(domain.EMAIL, email)

		ctx.Next()
	}
}
