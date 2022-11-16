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
		if strings.Contains(ctx.FullPath(), "docs") {
			ctx.Next()
			return
		}

		authToken := ctx.Request.Header.Get("Authorization")
		if authToken == "" {
			handler.ErrorResponse(ctx, http.StatusBadRequest, util.ErrNoAuthorization)
			ctx.Abort()
			return
		}

		authToken = strings.Replace(authToken, "Bearer ", "", 1)

		user_id, err := jwt.ValidateToken(authToken)
		if err != nil || user_id == "" {
			handler.ErrorResponse(ctx, http.StatusBadRequest, util.ErrTokenInvalid)
			ctx.Abort()
			return
		}

		ctx.Set(domain.EMAIL, user_id)

		ctx.Next()
	}
}
